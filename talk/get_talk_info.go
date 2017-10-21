package talk

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/devopsdays/devopsdays-cli/helpers/paths"
	"github.com/gernest/front"
	"github.com/pkg/errors"
)

// GetTalkInfo loads in a talk file and returns a struct with the information, decoded
// from the TOML in the frontmatter
func GetTalkInfo(file, city, year string) (talk *Talk, err error) {

	filePath := filepath.Join(paths.EventContentPath(city, year), "program", file)
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return talk, errors.Wrap(err, "load files failed")
	}
	m := front.NewMatter()
	m.Handle("+++", TOMLHandler)

	f, body, err := m.Parse(strings.NewReader(string(dat)))
	if err != nil {
		return talk, errors.Wrap(err, "get list of talks failed")
	}

	talk = &Talk{
		Name:         file,
		Title:        f["Title"].(string),
		Description:  f["Description"].(string),
		Speakers:     f["Speakers"].([]string),
		YouTube:      f["YouTube"].(string),
		Vimeo:        f["Vimeo"].(string),
		Speakerdeck:  f["Speakerdeck"].(string),
		Slideshare:   f["Slideshare"].(string),
		Googleslides: f["Googleslides"].(string),
		PDF:          f["PDF"].(string),
		Slides:       f["Slides"].(string),
		Abstract:     body,
	}

	return
}

// TOMLHandler decodes TOML string into a go map[string]interface{}
func TOMLHandler(front string) (map[string]interface{}, error) {

	var thisTalk Talk
	if _, err := toml.Decode(front, &thisTalk); err != nil {

		log.Fatal(err)
	}
	x := map[string]interface{}{
		"Name":         thisTalk.Name,
		"Title":        thisTalk.Title,
		"Description":  thisTalk.Description,
		"Speakers":     thisTalk.Speakers,
		"YouTube":      thisTalk.YouTube,
		"Vimeo":        thisTalk.Vimeo,
		"Speakerdeck":  thisTalk.Speakerdeck,
		"Slideshare":   thisTalk.Slideshare,
		"Googleslides": thisTalk.Googleslides,
		"PDF":          thisTalk.PDF,
		"Slides":       thisTalk.Slides,
		"Abstract":     thisTalk.Abstract,
	}

	return x, nil

}
