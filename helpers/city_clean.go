package helpers

import "strings"

//CityClean returns a city name all in lower case with spaces converted to dashes
func CityClean(city string) (cityClean string) {
	cityClean = strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)
	return
}
