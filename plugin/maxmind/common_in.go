package maxmind

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/xtls/geoip/lib"
)

var (
	defaultGeoLite2CountryMMDBFile = filepath.Join("./", "geolite2", "GeoLite2-Country.mmdb")
	defaultDBIPCountryMMDBFile     = filepath.Join("./", "db-ip", "dbip-country-lite.mmdb")
	defaultIPInfoLiteMMDBFile      = filepath.Join("./", "ipinfo", "ipinfo_lite.mmdb")
)

func defaultIPInfoLiteMMDBURI(uri string) string {
	if uri != "" {
		return uri
	}

	return defaultIPInfoLiteMMDBFile
}

func newGeoLite2CountryMMDBIn(iType string, iDesc string, action lib.Action, data json.RawMessage) (lib.InputConverter, error) {
	var tmp struct {
		URI                 string     `json:"uri"`
		Want                []string   `json:"wantedList"`
		OnlyIPType          lib.IPType `json:"onlyIPType"`
		IncludeAllCountries bool       `json:"includeAllCountries"`
	}

	if len(data) > 0 {
		if err := json.Unmarshal(data, &tmp); err != nil {
			return nil, fmt.Errorf("❌ [type %s | action %s] %w", iType, action, err)
		}
	}

	if tmp.URI == "" {
		switch iType {
		case TypeGeoLite2CountryMMDBIn:
			tmp.URI = defaultGeoLite2CountryMMDBFile

		case TypeDBIPCountryMMDBIn:
			tmp.URI = defaultDBIPCountryMMDBFile

		case TypeIPInfoCountryMMDBIn:
			tmp.URI = defaultIPInfoLiteMMDBURI(tmp.URI)
		}
	}

	// Filter want list
	wantList := make(map[string]bool)
	for _, want := range tmp.Want {
		if want = strings.ToUpper(strings.TrimSpace(want)); want != "" {
			wantList[want] = true
		}
	}

	return &GeoLite2CountryMMDBIn{
		Type:                iType,
		Action:              action,
		Description:         iDesc,
		URI:                 tmp.URI,
		Want:                wantList,
		OnlyIPType:          tmp.OnlyIPType,
		IncludeAllCountries: tmp.IncludeAllCountries,
	}, nil
}
