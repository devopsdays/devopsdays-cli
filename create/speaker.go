// Package create provides functions to create new content.
package create

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"text/template"

	"github.com/devopsdays/devopsdays-cli/helpers"
	"github.com/devopsdays/devopsdays-cli/model"
)

const tmpl = `+++
Title = "{{ .Title }}"
type = "speaker"
{{ with .Website }}website = "{{ . }}"{{ end }}
{{ with .Twitter }}twitter = "{{ . }}"{{ end }}
{{ with .Facebook }}facebook = "{{ . }}"{{ end }}
{{ with .Linkedin }}linkedin = "{{ . }}"{{ end }}
{{ with .Github }}github = "{{ . }}"{{ end }}
{{ with .Gitlab }}gitlab = "{{ . }}"{{ end }}
{{ with .ImagePath }}image = "{{ . }}"{{ end }}
+++
Food-truck SpaceTeam pivot earned media agile big data entrepreneur actionable insight iterate unicorn convergence driven moleskine. User centered design piverate physical computing disrupt moleskine co-working fund pivot. Waterfall is so 2000 and late integrate responsive big data agile piverate affordances. Agile earned media pivot viral engaging thought leader prototype workflow.

`

// NewSpeaker takes in a constructed Speaker type and generates the stuff
func NewSpeaker(speaker model.Speaker, city string, year string) (err error) {

	cleanName := helpers.NameClean(speaker.Name)
	t := template.New("Speaker template")

	t, err = t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	err = t.Execute(os.Stdout, speaker)
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}

	s := []string{strings.TrimSpace(cleanName), ".md"}
	f, err := os.Create(filepath.Join(helpers.EventContentPath(city, year), "speakers", strings.Join(s, "")))
	if err != nil {
		return err
	}
	defer f.Close()
	t.Execute(f, speaker)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created speaker file for", speaker.Title, "at", filepath.Join(helpers.EventContentPath(city, year), "speakers", strings.Join(s, "")))
	}
	return
}
