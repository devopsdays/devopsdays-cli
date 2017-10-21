package speaker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"text/template"

	"github.com/devopsdays/devopsdays-cli/helpers/paths"
	"github.com/devopsdays/devopsdays-cli/names"
	"github.com/fatih/color"
)

// NewSpeaker takes in a constructed Speaker type and generates the stuff
func NewSpeaker(speaker Speaker, city string, year string) (err error) {

	cleanName := names.NameClean(speaker.Name)
	t := template.New("Speaker template")

	t, err = t.Parse(speakerTmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	if err := os.MkdirAll(filepath.Join(paths.EventContentPath(city, year), "speakers"), 0777); err != nil {
		log.Fatal(err)
	}
	s := []string{strings.TrimSpace(cleanName), ".md"}
	f, err := os.Create(filepath.Join(paths.EventContentPath(city, year), "speakers", strings.Join(s, "")))
	if err != nil {
		return err
	}
	defer f.Close()
	t.Execute(f, speaker)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(color.Output, "\n\n\nCreated speaker file for %s\n", color.GreenString(speaker.Title))
		fmt.Fprintf(color.Output, "at %s\n\n\n", color.BlueString(filepath.Join(paths.EventContentPath(city, year), "speakers", strings.Join(s, ""))))
	}
	return
}
