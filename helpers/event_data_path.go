package helpers

import (
	"path/filepath"
	"strings"
)

// EventDataPath returns the full path the the data directory for events
func EventDataPath(webdir, city, year string) (eventDataPath string) {
	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10), ".yml"}
	eventDataPath = filepath.Join(webdir, "data", "events", strings.Join(s, ""))
	// eventDataPath = strings.Join(s, "")
	// eventDataPath = webdir
	return eventDataPath
}
