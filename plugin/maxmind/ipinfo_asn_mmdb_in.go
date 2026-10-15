package maxmind

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/oschwald/maxminddb-golang/v2"
	"github.com/xtls/geoip/lib"
)

const (
	TypeIPInfoASNMMDBIn = "ipinfoASNMMDB"
	DescIPInfoASNMMDBIn = "Convert IPInfo Lite ASN mmdb database to other formats"
)

func init() {
	lib.RegisterInputConfigCreator(TypeIPInfoASNMMDBIn, func(action lib.Action, data json.RawMessage) (lib.InputConverter, error) {
		return newIPInfoASNMMDBIn(action, data)
	})
	lib.RegisterInputConverter(TypeIPInfoASNMMDBIn, &IPInfoASNMMDBIn{
		Description: DescIPInfoASNMMDBIn,
	})
}

func newIPInfoASNMMDBIn(action lib.Action, data json.RawMessage) (lib.InputConverter, error) {
	var tmp struct {
		URI        string                 `json:"uri"`
		Want       lib.WantedListExtended `json:"wantedList"`
		OnlyIPType lib.IPType             `json:"onlyIPType"`
	}

	if len(data) > 0 {
		if err := json.Unmarshal(data, &tmp); err != nil {
			return nil, err
		}
	}

	tmp.URI = defaultIPInfoLiteMMDBURI(tmp.URI)

	return &IPInfoASNMMDBIn{
		Type:        TypeIPInfoASNMMDBIn,
		Action:      action,
		Description: DescIPInfoASNMMDBIn,
		URI:         tmp.URI,
		Want:        newIPInfoASNMMDBWantList(tmp.Want),
		OnlyIPType:  tmp.OnlyIPType,
	}, nil
}

type IPInfoASNMMDBIn struct {
	Type        string
	Action      lib.Action
	Description string
	URI         string
	Want        map[string][]string
	OnlyIPType  lib.IPType
}

func (i *IPInfoASNMMDBIn) GetType() string {
	return i.Type
}

func (i *IPInfoASNMMDBIn) GetAction() lib.Action {
	return i.Action
}

func (i *IPInfoASNMMDBIn) GetDescription() string {
	return i.Description
}

func (i *IPInfoASNMMDBIn) Input(container lib.Container) (lib.Container, error) {
	var content []byte
	var err error
	switch {
	case strings.HasPrefix(strings.ToLower(i.URI), "http://"), strings.HasPrefix(strings.ToLower(i.URI), "https://"):
		content, err = lib.GetRemoteURLContent(i.URI)
	default:
		content, err = os.ReadFile(i.URI)
	}
	if err != nil {
		return nil, err
	}

	entries := make(map[string]*lib.Entry)
	if err := i.generateEntries(content, entries); err != nil {
		return nil, err
	}

	if len(entries) == 0 {
		return nil, fmt.Errorf("[type %s | action %s] no entry is generated", i.Type, i.Action)
	}

	ignoreIPType := lib.GetIgnoreIPType(i.OnlyIPType)

	for _, entry := range entries {
		switch i.Action {
		case lib.ActionAdd:
			if err := container.Add(entry, ignoreIPType); err != nil {
				return nil, err
			}
		case lib.ActionRemove:
			if err := container.Remove(entry, lib.CaseRemovePrefix, ignoreIPType); err != nil {
				return nil, err
			}
		default:
			return nil, lib.ErrUnknownAction
		}
	}

	return container, nil
}

func (i *IPInfoASNMMDBIn) generateEntries(content []byte, entries map[string]*lib.Entry) error {
	db, err := maxminddb.OpenBytes(content)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Verify(); err != nil {
		return err
	}

	for network := range db.Networks() {
		var record ipInfoLite
		if err := network.Decode(&record); err != nil {
			return err
		}

		asn := normalizeIPInfoASN(record.ASN)
		if asn == "" || !network.Found() {
			continue
		}

		listNames := []string{"AS" + asn}
		if len(i.Want) > 0 {
			var found bool
			listNames, found = i.Want[asn]
			if !found {
				continue
			}
		}

		for _, listName := range listNames {
			entry, found := entries[listName]
			if !found {
				entry = lib.NewEntry(listName)
			}

			if err := entry.AddPrefix(network.Prefix()); err != nil {
				return err
			}

			entries[listName] = entry
		}
	}

	return nil
}

func newIPInfoASNMMDBWantList(want lib.WantedListExtended) map[string][]string {
	wantList := make(map[string][]string)

	for list, asnList := range want.TypeMap {
		list = strings.ToUpper(strings.TrimSpace(list))
		if list == "" {
			continue
		}

		for _, asn := range asnList {
			asn = normalizeIPInfoASN(asn)
			if asn == "" {
				continue
			}

			wantList[asn] = append(wantList[asn], list)
		}
	}

	for _, asn := range want.TypeSlice {
		asn = normalizeIPInfoASN(asn)
		if asn == "" {
			continue
		}

		wantList[asn] = []string{"AS" + asn}
	}

	return wantList
}

func normalizeIPInfoASN(asn string) string {
	return strings.TrimPrefix(strings.ToUpper(strings.TrimSpace(asn)), "AS")
}
