package maxmind

import (
	"bytes"
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
	"github.com/oschwald/maxminddb-golang/v2"
	"github.com/xtls/geoip/lib"
)

func TestMMDBRejectsCorruptSeparator(t *testing.T) {
	writer, err := mmdbwriter.New(mmdbwriter.Options{
		DatabaseType: "IPinfo-Test", IPVersion: 4, IncludeReservedNetworks: true,
		Description: map[string]string{"en": "IPInfo test fixture"}, Languages: []string{"en"},
	})
	if err != nil {
		t.Fatal(err)
	}
	_, network, err := net.ParseCIDR("192.0.2.0/24")
	if err != nil {
		t.Fatal(err)
	}
	if err := writer.Insert(network, mmdbtype.Map{
		"country_code": mmdbtype.String("US"), "asn": mmdbtype.String("AS64500"),
	}); err != nil {
		t.Fatal(err)
	}
	var buffer bytes.Buffer
	if _, err := writer.WriteTo(&buffer); err != nil {
		t.Fatal(err)
	}
	data := buffer.Bytes()
	db, err := maxminddb.OpenBytes(data)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Verify(); err != nil {
		t.Fatalf("valid fixture failed verification: %v", err)
	}
	t.Run("valid ASN input", func(t *testing.T) {
		entries := make(map[string]*lib.Entry)
		input := &IPInfoASNMMDBIn{Type: TypeIPInfoASNMMDBIn}
		if err := input.generateEntries(data, entries); err != nil {
			t.Fatal(err)
		}
		entry, found := entries["AS64500"]
		if !found {
			t.Fatal("valid ASN category was lost")
		}
		prefixes, err := entry.MarshalText()
		if err != nil || len(prefixes) != 1 || prefixes[0] != "192.0.2.0/24" {
			t.Fatalf("valid ASN prefixes = %v, %v", prefixes, err)
		}
	})
	separatorOffset := int(db.Metadata.NodeCount) * int(db.Metadata.RecordSize) / 4
	db.Close()
	data[separatorOffset] = 1
	db, err = maxminddb.OpenBytes(data)
	if err != nil {
		t.Fatalf("fixture must pass OpenBytes to exercise verification: %v", err)
	}
	verificationErr := db.Verify()
	db.Close()
	if verificationErr == nil {
		t.Fatal("corrupt fixture passed verification")
	}
	file := filepath.Join(t.TempDir(), "corrupt.mmdb")
	if err := os.WriteFile(file, data, 0600); err != nil {
		t.Fatal(err)
	}
	for _, kind := range []string{TypeGeoLite2CountryMMDBIn, TypeDBIPCountryMMDBIn, TypeIPInfoCountryMMDBIn, TypeIPInfoASNMMDBIn} {
		t.Run(kind, func(t *testing.T) {
			entries := make(map[string]*lib.Entry)
			var err error
			if kind == TypeIPInfoASNMMDBIn {
				err = (&IPInfoASNMMDBIn{Type: kind}).generateEntries(data, entries)
			} else {
				err = (&GeoLite2CountryMMDBIn{Type: kind}).generateEntries(data, entries)
			}
			if err == nil || err.Error() != verificationErr.Error() || len(entries) != 0 {
				t.Fatalf("corrupt input generated %d categories, error %v; want %v", len(entries), err, verificationErr)
			}
		})
	}
	t.Run("output source", func(t *testing.T) {
		output := &GeoLite2CountryMMDBOut{Type: TypeIPInfoCountryMMDBOut, SourceMMDBURI: file}
		info, err := output.GetExtraInfo()
		if err == nil || err.Error() != verificationErr.Error() || len(info) != 0 {
			t.Fatalf("corrupt output source returned %v, %v; want %v", info, err, verificationErr)
		}
	})
}
