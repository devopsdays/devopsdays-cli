package helpers

import "strings"

// NameClean returns a person's full name all in lower case with spaces converted to dashes
func NameClean(name string) (nameClean string) {
	nameClean = strings.Replace(strings.TrimSpace(strings.ToLower(name)), " ", "-", 10)
	return
}
