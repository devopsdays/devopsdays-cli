package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

// CheckSpeaker takes in three arguments, the city, the year, and the cleaned name of a speaker, and returns true if the speaker already exists.
func CheckSpeaker(city, year, speaker string) bool {
	s := []string{strings.TrimSpace(speaker), ".md"}
	if _, err := os.Stat(filepath.Join(EventContentPath(city, year), "speakers", strings.Join(s, ""))); err == nil {
		return true
	}
	return false

}
