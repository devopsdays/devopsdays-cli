package event

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/alecthomas/template"
	"github.com/devopsdays/devopsdays-cli/helpers/paths"
)

func createEventContentFile(city, year, page string) (string, error) {

	err := os.MkdirAll((paths.EventContentPath(city, year)), 0755)
	if err != nil {
		return "", err
	}

	// find a rice.Box
	// to compile, cd to event directory and run `rice embed-go`
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		log.Fatal(err)
	}
	templateName := page + ".md.tmpl"
	// get file contents as string
	templateString, err := templateBox.String(templateName)
	if err != nil {
		log.Fatal(err)
	}
	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)}
	slug := strings.Join(s, "")
	t, err := template.New(page+".md").Delims("[[", "]]").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		City      string
		Year      string
		Slug      string
		CityClean string
	}{
		strings.TrimSpace(city),
		strings.TrimSpace(year),
		slug,
		CityClean(city),
	}
	filePath := filepath.Join((paths.EventContentPath(city, year)), (page + ".md"))
	f, err := os.Create(filePath)
	if err != nil {
		return "Cannot create", err
	}
	defer f.Close()
	t.Execute(f, data)
	if err != nil {
		fmt.Println(err, "template execute error")
	} else {
		fmt.Println("Created event content file for", city, "for year", year, "at", filePath)
	}
	return city, nil

}
