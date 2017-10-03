package helpers

import (
	"testing"
)

func TestNameClean(t *testing.T) {
	t.Log("Cleaning name value... (expected result: 'george-bluth')")
	cleanName := NameClean("George Bluth")

	if cleanName != "george-bluth" {
		t.Errorf("Expected result of 'george-bluth' but it was %s instead.", cleanName)
	}
}
