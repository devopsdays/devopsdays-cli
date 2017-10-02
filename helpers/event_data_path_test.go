package helpers

import (
	"testing"
)

func TestEventDataPath(t *testing.T) {
	t.Log("Testing eventDataPath function... (expected result: '" + GetWebdir() + "/data/events/2018-new-york.yml')")
	testDataPath := EventDataPath(GetWebdir(), "New York", "2018")
	if testDataPath != GetWebdir()+"/data/events/2018-new-york.yml" {
		t.Errorf("Expected result of '"+GetWebdir()+"/data/events/2018-new-york.yml' but it was %s instead.", testDataPath)
	}
}
