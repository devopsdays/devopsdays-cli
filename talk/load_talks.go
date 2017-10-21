package talk

import (
	"io/ioutil"
	"path/filepath"

	"github.com/devopsdays/devopsdays-cli/helpers/paths"
	"github.com/pkg/errors"
)

func loadTalks(city, year string) (talks []*Talk, err error) {
	talksDir := filepath.Join(paths.EventContentPath(city, year), "program")
	files, err := ioutil.ReadDir(talksDir)
	if err != nil {
		return nil, errors.Wrap(err, "read directory failed")
	}

	for _, f := range files {
		thisTalk, _ := GetTalkInfo(f.Name(), city, year)
		talks = append(talks, thisTalk)
	}

	return
}
