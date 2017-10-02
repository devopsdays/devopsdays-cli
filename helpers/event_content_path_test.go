package helpers

import (
	"testing"
)

func TestEventContentPath(t *testing.T) {
	t.Log("Testing eventContentPath function... (expected result: '/Users/mattstratton/src/devopsdays-web/content/events/2018-new-york')")
	testContentPath := EventContentPath("New York", "2018")
	if testContentPath != GetWebdir()+"/content/events/2018-new-york" {
		t.Errorf("Expected result of '/Users/mattstratton/src/devopsdays-web/content/events/2018-new-york' but it was %s instead.", testContentPath)
	}
}
