package maxmind

import (
	"bytes"
	"encoding/json"
	"errors"
	"net"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
	"github.com/xtls/geoip/lib"
)

func TestCountryMMDBIncludesMissingCountries(t *testing.T) {
	for _, inputType := range []string{TypeGeoLite2CountryMMDBIn, TypeDBIPCountryMMDBIn, TypeIPInfoCountryMMDBIn} {
		writer, err := mmdbwriter.New(mmdbwriter.Options{
			DatabaseType: "Country-Test", IPVersion: 4, IncludeReservedNetworks: true,
			Description: map[string]string{"en": "Country test fixture"}, Languages: []string{"en"},
		})
		if err != nil {
			t.Fatal(err)
		}
		_, network, err := net.ParseCIDR("192.0.2.0/24")
		if err != nil {
			t.Fatal(err)
		}
		record := mmdbtype.Map{"country": mmdbtype.Map{"iso_code": mmdbtype.String("US")}}
		if inputType == TypeIPInfoCountryMMDBIn {
			record = mmdbtype.Map{"country_code": mmdbtype.String("US")}
		}
		if err := writer.Insert(network, record); err != nil {
			t.Fatal(err)
		}
		var data bytes.Buffer
		if _, err := writer.WriteTo(&data); err != nil {
			t.Fatal(err)
		}
		file := filepath.Join(t.TempDir(), "country.mmdb")
		if err := os.WriteFile(file, data.Bytes(), 0600); err != nil {
			t.Fatal(err)
		}
		for _, tc := range []struct {
			name       string
			includeAll bool
			wanted     []string
			onlyIPType lib.IPType
			action     lib.Action
			names      []string
		}{
			{name: "include missing wanted country", includeAll: true, wanted: []string{" cn ", "us"}, names: []string{"CN", "US"}},
			{name: "only countries with data", wanted: []string{"cn", "us"}, names: []string{"US"}},
			{name: "respect wanted list", includeAll: true, wanted: []string{"us"}, names: []string{"US"}},
			{name: "family filter retains empty categories", includeAll: true, wanted: []string{"cn", "us"}, onlyIPType: lib.IPv6, names: []string{"CN", "US"}},
			{name: "removal does not create missing countries", includeAll: true, wanted: []string{"cn", "us"}, action: lib.ActionRemove, names: []string{"US"}},
		} {
			t.Run(inputType+"/"+tc.name, func(t *testing.T) {
				action := tc.action
				if action == "" {
					action = lib.ActionAdd
				}
				container := lib.NewContainer()
				if action == lib.ActionRemove {
					entry := lib.NewEntry("US")
					if err := entry.AddPrefix("192.0.2.0/24"); err != nil {
						t.Fatal(err)
					}
					if err := container.Add(entry); err != nil {
						t.Fatal(err)
					}
				}
				config, err := json.Marshal(map[string]any{"input": []any{map[string]any{
					"type": inputType, "action": action, "args": map[string]any{
						"uri": file, "includeAllCountries": tc.includeAll,
						"wantedList": tc.wanted, "onlyIPType": tc.onlyIPType,
					},
				}}})
				if err != nil {
					t.Fatal(err)
				}
				instance, err := lib.NewInstance()
				if err != nil {
					t.Fatal(err)
				}
				if err := instance.InitConfigFromBytes(config); err != nil {
					t.Fatal(err)
				}
				if err := instance.RunInput(container); err != nil {
					t.Fatal(err)
				}
				var names []string
				for entry := range container.Loop() {
					names = append(names, entry.GetName())
					got, err := entry.MarshalText()
					if entry.GetName() == "CN" || tc.onlyIPType == lib.IPv6 || action == lib.ActionRemove {
						if !errors.Is(err, lib.ErrEmptyPrefix) {
							t.Fatalf("%s should be empty: %v, %v", entry.GetName(), got, err)
						}
					} else if err != nil || !slices.Equal(got, []string{"192.0.2.0/24"}) {
						t.Fatalf("US prefixes = %v, %v", got, err)
					}
				}
				slices.Sort(names)
				if !slices.Equal(names, tc.names) {
					t.Fatalf("country categories = %v; want %v", names, tc.names)
				}
			})
		}
	}
}
