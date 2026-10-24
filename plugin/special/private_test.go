package special

import (
	"testing"

	"github.com/xtls/geoip/lib"
)

func TestPrivateSpecialPurposeSelection(t *testing.T) {
	input, err := newPrivate(lib.ActionAdd, nil)
	if err != nil {
		t.Fatal(err)
	}
	container, err := input.Input(lib.NewContainer())
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		query string
		want  bool
	}{
		// Keep assigned services without matching their entire parent blocks.
		{"192.0.0.0/29", true},
		{"192.0.0.8/31", true},
		{"192.0.0.10", true},
		{"192.0.0.8/30", false},
		{"192.0.0.170/31", true},
		{"192.0.0.168/30", false},
		{"192.0.0.255", false},
		{"192.52.193.0/24", true},
		{"192.88.99.2", true},
		{"192.88.99.1", false},
		{"192.88.99.3", false},
		{"192.88.99.0/24", false},
		{"2001::/32", true},
		{"2001:1::1", true},
		{"2001:1::2", true},
		{"2001:1::3", true},
		{"2001:1::4", false},
		{"2001:2::/48", true},
		{"2001:3::/32", true},
		{"2001:20::/28", true},
		{"2001:30::/28", true},
		{"2001:100::1", false},
		{"2001::/23", false},

		// Retain selected translation and special-purpose allocations.
		{"64:ff9b::/96", true},
		{"64:ff9b:1::/48", true},
		{"100:0:0:1::/64", true},
		{"2002::/16", true},
		{"3fff::/20", true},
		{"5f00::/16", true},

		// AS112, deprecated ORCHID, and ordinary public addresses stay excluded.
		{"192.31.196.0/24", false},
		{"192.175.48.0/24", false},
		{"2001:4:112::/48", false},
		{"2001:10::/28", false},
		{"2620:4f:8000::/48", false},
		{"8.8.8.8", false},
		{"::ffff:8.8.8.8", false},
		{"2001:4860:4860::8888", false},
	} {
		t.Run(tc.query, func(t *testing.T) {
			_, got, err := container.Lookup(tc.query, entryNamePrivate)
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.want {
				t.Fatalf("private contains %s = %v; want %v", tc.query, got, tc.want)
			}
		})
	}
}
