package main

import (
	"flag"
	"path/filepath"
	"testing"

	"github.com/ProfoundNetworks/gpntest"
)

var binary_name string = "domainify"
var update *bool

func init() {
	testing.Init()
	update = flag.Bool("u", false, "update .golden files")
	flag.Parse()
}

func TestDomainifyArgs(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		args string
		want string
	}{
		{"www.profound.net", "profound.net\n"},
		{"www.futurium.ec.europa.eu home.dotdashmeredith.mediaroom.com",
			"futurium.ec.europa.eu\ndotdashmeredith.mediaroom.com\n"},
		{"www.openfusion.com.au com", "openfusion.com.au\n\n"},
	}

	cmd0 := gpntest.FindBinary(t, binary_name)
	for _, tc := range tests {
		cmd := cmd0 + " " + tc.args

		got := gpntest.MustExecuteCommand(t, cmd)

		if string(got) != tc.want {
			t.Errorf("want %q, got %q", tc.want, got)
		}
	}
}

func TestDomainStdin(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		infile string
	}{
		{"test1.txt"},
	}

	cmd0 := gpntest.FindBinary(t, binary_name)
	for _, tc := range tests {
		cmd := cmd0 + " --stdin <" +
			filepath.Join("testdata", tc.infile)

		got := gpntest.MustExecuteCommand(t, cmd)

		gpntest.UpdateGolden(t, update, tc.infile, got)

		gpntest.CompareFiles(t, tc.infile, tc.infile, got)
	}
}
