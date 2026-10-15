package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestLookupDirectoryFlag(t *testing.T) {
	// The command uses log.Fatal, so exercise its validation in a child process.
	if format := os.Getenv("GEOIP_TEST_LOOKUP_FORMAT"); format != "" {
		rootCmd.SetArgs([]string{"lookup", "--format", format, "--dir", os.Getenv("GEOIP_TEST_LOOKUP_DIR"), "192.0.2.1"})
		main()
		return
	}
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "test.txt"), []byte("192.0.2.0/24\n"), 0600); err != nil {
		t.Fatal(err)
	}
	for _, format := range []string{"xrayGeoIPDat", "maxmindMMDB", "dbipCountryMMDB", "ipinfoCountryMMDB", "ipinfoASNMMDB", "text"} {
		t.Run(format, func(t *testing.T) {
			cmd := exec.Command(os.Args[0], "-test.run=^TestLookupDirectoryFlag$")
			cmd.Env = append(os.Environ(), "GEOIP_TEST_LOOKUP_FORMAT="+format, "GEOIP_TEST_LOOKUP_DIR="+dir)
			output, err := cmd.CombinedOutput()
			if format == "text" {
				if err != nil || !strings.HasPrefix(string(output), "test\n") {
					t.Fatalf("text directory lookup = %s, %v", output, err)
				}
				return
			}
			if err == nil || !strings.Contains(string(output), "does not support \"dir\" flag") || !strings.Contains(string(output), "please use \"uri\" flag instead") {
				t.Fatalf("single-file directory lookup = %s, %v", output, err)
			}
		})
	}
}
