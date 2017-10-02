package helpers

import (
	"testing"
)

func TestCityClean(t *testing.T) {
	t.Log("Cleaning city value... (expected result: 'new-york')")
	cleanCity := CityClean("New York")

	if cleanCity != "new-york" {
		t.Errorf("Expected result of 'new-york' but it was %s instead.", cleanCity)
	}
}
