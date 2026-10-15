package xray

import (
	"bytes"
	"net/netip"
	"slices"
	"testing"

	"github.com/xtls/geoip/lib"
	"google.golang.org/protobuf/proto"
)

func TestDATInputCIDRParsing(t *testing.T) {
	ipv4 := netip.MustParseAddr("192.0.2.1")
	mapped := ipv4.As16()
	for _, tc := range []struct {
		name string
		ip   []byte
		bits uint32
		want string
	}{
		{"IPv4", ipv4.AsSlice(), 24, "192.0.2.0/24"},
		{"mapped with IPv4 mask", mapped[:], 24, "192.0.2.0/24"},
		{"mapped with IPv6 mask", mapped[:], 120, "192.0.2.0/24"},
		{"mapped default route", mapped[:], 96, "0.0.0.0/0"},
		{"mapped host", mapped[:], 128, "192.0.2.1/32"},
		{"IPv6", netip.MustParseAddr("2001:db8::1").AsSlice(), 32, "2001:db8::/32"},
		{"invalid bytes", []byte{1, 2, 3}, 24, ""},
		{"missing bytes", nil, 0, ""},
		{"invalid IPv4 mask", ipv4.AsSlice(), 33, ""},
		{"invalid mapped mask", mapped[:], 95, ""},
		{"invalid IPv6 mask", netip.MustParseAddr("2001:db8::1").AsSlice(), 129, ""},
		{"overflowing mask", ipv4.AsSlice(), ^uint32(0), ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			data, err := proto.Marshal(&GeoIPList{Entry: []*GeoIP{{
				CountryCode: " test ", Cidr: []*CIDR{{Ip: tc.ip, Prefix: tc.bits}},
			}}})
			if err != nil {
				t.Fatal(err)
			}
			input := &GeoIPDatIn{Type: TypeGeoIPDatIn, Action: lib.ActionAdd}
			entries := make(map[string]*lib.Entry)
			err = input.generateEntries(bytes.NewReader(data), entries)
			if tc.want == "" {
				if err == nil {
					t.Fatal("invalid DAT CIDR was accepted")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			entry, found := entries["TEST"]
			if !found {
				t.Fatal("missing normalized category")
			}
			got, err := entry.MarshalText()
			if err != nil || !slices.Equal(got, []string{tc.want}) {
				t.Fatalf("DAT prefixes = %v, %v; want %s", got, err, tc.want)
			}
		})
	}
}
