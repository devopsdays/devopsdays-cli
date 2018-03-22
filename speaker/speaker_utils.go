package speaker

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	paths "github.com/devopsdays/devopsdays-cli/helpers/paths"
	"github.com/gernest/front"
	"github.com/pkg/errors"
)

// GetSpeakers takes in the city and year and returns a list of the talks
func GetSpeakers(city, year string) ([]string, error) {

	speakerdir := filepath.Join(paths.EventContentPath(city, year), "speakers")

	files, err := ioutil.ReadDir(speakerdir)

	if err != nil {
		return nil, errors.Wrap(err, "read speaker directory failed")
	}
	var s []string
	for _, f := range files {
		s = append(s, f.Name())
	}
	return s, nil
}

// GetSpeakerInfo takes the filename, city, and year of a speaker and returns a Speaker object with the details.
func GetSpeakerInfo(file, city, year string) (speaker Speaker, err error) {

	filePath := filepath.Join(paths.EventContentPath(city, year), "speakers", file)
	dat, err := ioutil.ReadFile(filePath)
	m := front.NewMatter()
	m.Handle("+++", TOMLHandler)

	f, body, err := m.Parse(strings.NewReader(string(dat)))
	if err != nil {
		panic(err)
	}

	speaker = Speaker{
		Name:      file,
		Title:     f["Title"].(string),
		Website:   f["Website"].(string),
		Twitter:   f["Twitter"].(string),
		Facebook:  f["Facebook"].(string),
		Linkedin:  f["Linkedin"].(string),
		Github:    f["Github"].(string),
		Gitlab:    f["Gitlab"].(string),
		ImagePath: f["ImagePath"].(string),
		Bio:       body,
	}

	return
}

// TOMLHandler decodes TOML string into a go map[string]interface{}
func TOMLHandler(front string) (map[string]interface{}, error) {

	var stuff Speaker
	if _, err := toml.Decode(front, &stuff); err != nil {
		log.Fatal(err)
	}
	x := map[string]interface{}{
		"Name":      stuff.Name,
		"Title":     stuff.Title,
		"Website":   stuff.Website,
		"Twitter":   stuff.Twitter,
		"Facebook":  stuff.Facebook,
		"Linkedin":  stuff.Linkedin,
		"Github":    stuff.Github,
		"Gitlab":    stuff.Gitlab,
		"ImagePath": stuff.ImagePath,
		"Bio":       stuff.Bio,
	}

	return x, nil
}
