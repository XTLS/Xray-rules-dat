package plaintext

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/xtls/geoip/lib"
)

func TestTextInputDirNameWithAtSign(t *testing.T) {
	dir := t.TempDir()
	wantCIDRs := []string{"192.0.2.0/24", "2001:db8::/48"}
	for _, filename := range []string{"in-cn.txt", "in-cn@cTcC.txt"} {
		if err := os.WriteFile(filepath.Join(dir, filename), []byte("192.0.2.0/24\n2001:db8::/48\n"), 0600); err != nil {
			t.Fatal(err)
		}
	}

	for _, tt := range []struct {
		name   string
		wanted []string
		want   []string
	}{
		{name: "all", want: []string{"IN-CN", "IN-CN@CTCC"}},
		{name: "wanted category with at sign", wanted: []string{"in-cn@ctcc"}, want: []string{"IN-CN@CTCC"}},
		{name: "wanted base category", wanted: []string{"in-cn"}, want: []string{"IN-CN"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			args, err := json.Marshal(map[string]any{"inputDir": dir, "wantedList": tt.wanted})
			if err != nil {
				t.Fatal(err)
			}
			input, err := newTextIn(TypeTextIn, "", lib.ActionAdd, args)
			if err != nil {
				t.Fatal(err)
			}
			container, err := input.Input(lib.NewContainer())
			if err != nil {
				t.Fatal(err)
			}
			if container.Len() != len(tt.want) {
				t.Fatalf("got %d categories, want %v", container.Len(), tt.want)
			}
			for _, name := range tt.want {
				entry, found := container.GetEntry(name)
				if !found {
					t.Fatalf("missing category %s", name)
				}
				if entry.GetName() != name {
					t.Fatalf("got category name %q, want %q", entry.GetName(), name)
				}
				got, err := entry.MarshalText()
				if err != nil || !slices.Equal(got, wantCIDRs) {
					t.Fatalf("%s prefixes = %v, %v; want %v", name, got, err, wantCIDRs)
				}
			}
		})
	}
}
