package event

import (
	"os"

	"github.com/devopsdays/devopsdays-cli/helpers/paths"
)

// CheckEvent takes in two arguments, the city and the year, and returns true if the city  exists.
func CheckEvent(city, year string) bool {
	if _, err := os.Stat(paths.EventDataPath(paths.GetWebdir(), city, year)); err == nil {
		return true
	}
	return false

}
