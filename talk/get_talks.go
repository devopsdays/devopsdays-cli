package talk

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/devopsdays/devopsdays-cli/helpers/paths"
)

// GetTalks takes in the city and year and returns a list of the talks
func GetTalks(city, year string) []string {

	talksDir := filepath.Join(paths.EventContentPath(city, year), "program")

	fmt.Println(talksDir)

	// @TODO implement error checking in helpers.GetTalks
	files, _ := ioutil.ReadDir(talksDir)
	// if err != nil {
	// 	return nil, err
	// }

	s := make([]string, len(files))
	for _, f := range files {
		s = append(s, f.Name())
	}
	return s
}
