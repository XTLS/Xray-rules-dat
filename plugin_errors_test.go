package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"github.com/xtls/geoip/lib"
)

func TestInputErrorsPreserveCauseAndContext(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing")
	for _, kind := range []string{"text", "xrayGeoIPDat", "mihomoMRS", "singboxSRS", "maxmindMMDB", "dbipCountryMMDB", "ipinfoCountryMMDB", "ipinfoASNMMDB", "maxmindGeoLite2CountryCSV", "maxmindGeoLite2ASNCSV"} {
		t.Run(kind, func(t *testing.T) {
			data, err := json.Marshal(map[string]any{"input": []any{map[string]any{
				"type": kind, "action": lib.ActionAdd, "args": map[string]any{
					"name": "TEST", "uri": missing, "country": missing, "ipv4": missing,
				},
			}}})
			if err != nil {
				t.Fatal(err)
			}
			instance, err := lib.NewInstance()
			if err != nil {
				t.Fatal(err)
			}
			if err := instance.InitConfigFromBytes(data); err != nil {
				t.Fatal(err)
			}
			err = instance.RunInput(lib.NewContainer())
			if !errors.Is(err, fs.ErrNotExist) {
				t.Fatalf("input error lost its file error cause: %v", err)
			}
			prefix := fmt.Sprintf("[type %s | action %s]", kind, lib.ActionAdd)
			if strings.Count(err.Error(), prefix) != 1 {
				t.Fatalf("input error must contain context exactly once: %v", err)
			}
		})
	}
}

func TestConfigErrorsPreserveCauseAndContext(t *testing.T) {
	for _, direction := range []string{"input", "output"} {
		for _, kind := range []string{"text", "xrayGeoIPDat", "mihomoMRS", "singboxSRS", "maxmindMMDB", "ipinfoCountryMMDB"} {
			t.Run(direction+"/"+kind, func(t *testing.T) {
				action := lib.ActionAdd
				if direction == "output" {
					action = lib.ActionOutput
				}
				data, err := json.Marshal(map[string]any{direction: []any{map[string]any{
					"type": kind, "action": action, "args": map[string]any{"onlyIPType": false},
				}}})
				if err != nil {
					t.Fatal(err)
				}
				instance, err := lib.NewInstance()
				if err != nil {
					t.Fatal(err)
				}
				err = instance.InitConfigFromBytes(data)
				var typeErr *json.UnmarshalTypeError
				if !errors.As(err, &typeErr) {
					t.Fatalf("configuration error lost its JSON cause: %v", err)
				}
				prefix := fmt.Sprintf("[type %s | action %s]", kind, action)
				if strings.Count(err.Error(), prefix) != 1 {
					t.Fatalf("configuration error must contain context exactly once: %v", err)
				}
			})
		}
	}
}
