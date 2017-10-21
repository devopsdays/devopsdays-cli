package talk

import (
	"io/ioutil"
	"path/filepath"

	"github.com/devopsdays/devopsdays-cli/helpers/paths"
	"github.com/pkg/errors"
)

// GetTalks takes in the city and year and returns a list of the talks
func GetTalks(city, year string) ([]string, error) {
	talksDir := filepath.Join(paths.EventContentPath(city, year), "program")

	files, err := ioutil.ReadDir(talksDir)

	if err != nil {
		return nil, errors.Wrap(err, "read directory failed")
	}

	var s []string

	for _, f := range files {
		s = append(s, f.Name())
	}
	return s, nil
}
