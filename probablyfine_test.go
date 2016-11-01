package main

import "testing"

func TestCityClean(t *testing.T) {
	cleanCity := cityClean("New York")

	if cleanCity != "new-york" {
		t.Error("Response from cityClean is unexpected value")
	}
}

func TestEventDataPath(t *testing.T) {
	testDataPath := eventDataPath("New York", "2018")

	if testDataPath != "/Users/mattstratton/src/devopsdays-web/data/events/2018-new-york.yml" {
		t.Error("Response from eventDataPath is an unexpected value")
	}
}
