package helpers

import (
	"path/filepath"
	"strings"
)

// EventStaticPath returns the full path of the static directory for an event
func EventStaticPath(webdir, city, year string) (eventStaticPath string) {
	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10), ".yml"}
	return filepath.Join(webdir, "static", "events", strings.Join(s, ""))
}
