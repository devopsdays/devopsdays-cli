package main

import "testing"

func TestCityClean(t *testing.T) {
	t.Log("Cleaning city value... (expected result: 'new-york')")
	cleanCity := cityClean("New York")

	if cleanCity != "new-york" {
		t.Errorf("Expected result of 'new-york' but it was %s instead.", cleanCity)
	}
}

func TestEventDataPath(t *testing.T) {
	t.Log("Testing eventDataPath function... (expected result: '/Users/mattstratton/src/devopsdays-web/data/events/2018-new-york.yml')")
	testDataPath := eventDataPath("New York", "2018") // TODO: Pass in webdir path to function

	if testDataPath != "/Users/mattstratton/src/devopsdays-web/data/events/2018-new-york.yml" {
		t.Error("Response from eventDataPath is an unexpected value")
	}
}

func TestValidateField(t *testing.T) {
	t.Log("Testing validateField() function")
	if v := validateField("Chicago", "city"); v != true {
		t.Error("Valid city did not pass validation test in validateField")
	}
	if v := validateField("3yl0RmG1wU8q5TeDPKZEsNU3E54nyYf5MNhGhzqcxhoLJkeckXCa1saWCPM24YhwIteGEUjLW8S715WkoDvt3vFsMaVeYXCUZWNL", "city"); v == true {
		t.Error("Invalid city passed validation test in validateField")
	}
	if v := validateField("2017", "year"); v != true {
		t.Error("Valid year did not pass validation test in validateField")
	}
	if v := validateField("19008", "year"); v == true {
		t.Error("Invalid year passed validation test in validateField")
	}
	if v := validateField("2011", "year"); v == true {
		t.Error("Invalid year passed validation test in validateField")
	}
	if v := validateField("2031", "year"); v == true {
		t.Error("Invalid year passed validation test in validateField")
	}
	if v := validateField("devopsdays", "twitter"); v != true {
		t.Error("Valid twitter did not pass validation test in validateField")
	}
	if v := validateField("devops days", "twitter"); v == true {
		t.Error("Invalid twitter passed validation test in validateField")
	}
}

func TestSponsorDataPath(t *testing.T) {
	t.Log("Testing evaluation of sponsor data file path")
	if v := sponsorDataPath("/Users/mattstratton/src/devopsdays-web", "chef"); v != "/Users/mattstratton/src/devopsdays-web/data/sponsors/chef.yml" {
		t.Error("Response from eventDataPath is an unexpected value")
	}
}
