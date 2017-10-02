package helpers

import (
	"path/filepath"
	"strings"
)

// EventContentPath returns the path for content for an event based upon city and year
func EventContentPath(city, year string) (eventContentPath string) {
	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)}
	eventContentPath = filepath.Join(GetWebdir(), "content", "events", strings.Join(s, ""))
	// eventContentPath = webdir
	return eventContentPath
}
