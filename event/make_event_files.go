package event

import (
	"os"
	"path/filepath"
	"strings"

	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/devopsdays/devopsdays-cli/helpers/paths"
	"github.com/pkg/errors"
)

func createEventContentFile(city, year, page string) error {

	err := os.MkdirAll((paths.EventContentPath(city, year)), 0755)
	if err != nil {
		return errors.Wrap(err, "make event content directory failed")
	}

	// find a rice.Box
	// to compile, cd to event directory and run `rice embed-go`
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		return errors.Wrap(err, "content template find failed")
	}
	templateName := page + ".md.tmpl"
	// get file contents as string
	templateString, err := templateBox.String(templateName)
	if err != nil {
		return errors.Wrapf(err, "cannot load template for %s", templateName)
	}
	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)}
	slug := strings.Join(s, "")
	t, err := template.New(page+".md").Delims("[[", "]]").Parse(templateString)
	if err != nil {
		return errors.Wrapf(err, "template parse failed for %s", templateName)
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
		return errors.Wrapf(err, "cannot create file for %s", filePath)
	}
	defer f.Close()
	t.Execute(f, data)
	if err != nil {
		return errors.Wrapf(err, "template execute error for %s", templateName)
	}
	return nil

}
