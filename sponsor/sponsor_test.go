package sponsor

import "testing"

// func TestSponsorDataPath(t *testing.T) {
// 	t.Log("Testing sponsorDataPath function... (expected result: '" + paths.GetWebdir() + "/data/events/2018-new-york.yml')")
// 	testDataPath := sponsorDataPath("chef")
// 	if testDataPath != paths.GetWebDir()+"/data/sponsors/chef.yml" {
// 		t.Errorf("Expected result of '"+paths.GetWebDir()+"data/events/2018-new-york.yml' but it was %s instead.", testDataPath)
// 	}
// }

// func TestSponsorImagePath(t *testing.T) {
// 	t.Log("Testing sponsorImagePath function... (expected result: '" + webdir + "/static/img/sponsors/chef.png')")
// 	testImagePath := sponsorImagePath("chef")
// 	if testImagePath != webdir+"/static/img/sponsors/chef.png" {
// 		t.Errorf("Expected result of '"+webdir+"/static/img/sponsors/chef.png' but it was %s instead.", testImagePath)
// 	}
// }

func TestCheckSponsor(t *testing.T) {
	t.Log("Testing checkSponsor function... (expected result: 'true'")
	testCheckSponsor := checkSponsor("chef")
	if testCheckSponsor != true {
		t.Errorf("Expected result of 'true' but didn't get it")
	}
}

// func TestCheckSponsorImage(t *testing.T) {
// 	t.Log("Testing checkSponsorImage function... (expected result: 'true'")
// 	s := webdir + "/static/img/sponsors/chef.png"
// 	testCheckSponsor := checkSponsorImage(s)
// 	if testCheckSponsor != true {
// 		t.Errorf("Expected result of 'true' but didn't get it")
// 	}
// }
