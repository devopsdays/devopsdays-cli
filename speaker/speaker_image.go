package speaker

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/devopsdays/devopsdays-cli/helpers/paths"
	"github.com/devopsdays/devopsdays-cli/images"
)

// SpeakerImage takes in a path to an image and resizes it to the proper dimensions and copies it to the destination
func SpeakerImage(srcPath, speaker, city, year string) (imageFile string) {

	var eventStaticPath string
	var err error
	eventStaticPath, err = paths.EventStaticPath(city, year)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(filepath.Join(eventStaticPath, "speakers"), 0777); err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`\.[^.]+$`)
	ext := strings.ToLower(re.FindString(srcPath))
	switch ext {
	case ".jpg":
		s := []string{strings.TrimSpace(speaker), ".jpg"}
		destPath := filepath.Join(eventStaticPath, "speakers", strings.Join(s, ""))
		images.ResizeImage(srcPath, destPath, "jpg", 600, 600)
		return strings.Join(s, "")
	case ".jpeg":
		s := []string{strings.TrimSpace(speaker), ".jpg"}
		destPath := filepath.Join(eventStaticPath, "speakers", strings.Join(s, ""))
		images.ResizeImage(srcPath, destPath, "jpg", 600, 600)
		return strings.Join(s, "")
	case ".png":
		s := []string{strings.TrimSpace(speaker), ".png"}
		destPath := filepath.Join(eventStaticPath, "speakers", strings.Join(s, ""))
		images.ResizeImage(srcPath, destPath, "png", 600, 600)
		return strings.Join(s, "")
	}
	return "busted"
}
