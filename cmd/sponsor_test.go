package cmd

import "testing"

func TestSponsorDataPath(t *testing.T) {
	t.Log("Testing evaluation of sponsor data file path")
	if v := sponsorDataPath("/Users/mattstratton/src/devopsdays-web", "chef"); v != "/Users/mattstratton/src/devopsdays-web/data/sponsors/chef.yml" {
		t.Error("Response from eventDataPath is an unexpected value")
	}
}
