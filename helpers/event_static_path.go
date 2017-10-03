package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

// EventStaticPath returns the full path of the static directory for an event, creating it if needed
func EventStaticPath(city, year string) (eventStaticPath string, err error) {
	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)}
	eventStaticPath = filepath.Join(GetWebdir(), "static", "events", strings.Join(s, ""))
	if err := os.MkdirAll(eventStaticPath, 0777); err == nil {
		return eventStaticPath, err
	}
	return "", err
}
