package helpers

import "os"

// CheckEvent takes in two arguments, the city and the year, and returns true if the city  exists.
func CheckEvent(city, year string) bool {
	if _, err := os.Stat(EventDataPath(GetWebdir(), city, year)); err == nil {
		return true
	}
	return false

}
