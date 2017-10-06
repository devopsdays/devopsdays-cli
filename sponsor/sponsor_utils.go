package sponsor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/devopsdays/devopsdays-cli/helpers/paths"
)

// checkSponsor takes in one argument, the name of a sponsor, and returns true if the sponsor already exists.
func checkSponsor(sponsor string) bool {
	if _, err := os.Stat(sponsorDataPath(sponsor)); err == nil {
		return true
	}
	return false

}

func checkSponsorImage(path string) bool {
	fmt.Println(path)
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false

}

func sponsorDataPath(sponsor string) (sponsorDataPath string) {
	s := []string{strings.TrimSpace(sponsor), ".yml"}
	sponsorDataPath = filepath.Join(paths.GetWebdir(), "data", "sponsors", strings.Join(s, ""))
	return sponsorDataPath
}

func sponsorImagePath(sponsor string) (sponsorImagePath string) {
	s := []string{paths.GetWebdir(), "/static/img/sponsors/", strings.TrimSpace(sponsor), ".png"}
	sponsorImagePath = strings.Join(s, "")
	return sponsorImagePath
}
