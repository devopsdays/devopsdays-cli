package commands

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

var (
	// Version is the current version of the devopsdays-cli tool. Unless set elsehwere, it is referred to as "master"
	Version = "master"
	// Build is the current build of the devopsdays-cli tool.
	// @todo The Build variable in cmd/version.go needs to be set somewhere.
	Build string
)

func init() {
	showCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of devopsdays-cli",
	Long:  `All software has versions. This is ours.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("devopsdays-cli version: ", Version)
		fmt.Println("devopsdays-cli build: ", Build)
		fmt.Println("hugo version: ", getHugoVersion())
		fmt.Println("devopsdays-theme version: ", getThemeVersion())
	},
}

func getHugoVersion() (hugoVersion string) {
	out, err := exec.Command("hugo", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out[:])
	re := regexp.MustCompile("v....")
	hugoVersion = re.FindString(s)
	return
}

func getThemeVersion() (themeVersion string) {
	var theme Theme
	themePath := filepath.Join(webdir, "themes", "devopsdays-theme", "theme.toml")
	if _, err := toml.DecodeFile(themePath, &theme); err != nil {
		fmt.Println(err)
		return
	}
	themeVersion = theme.Version
	return

}

// Theme represents the currently installed devopsdays-theme Hugo theme.
// The field Version represents the current version.
type Theme struct {
	Version string `toml:"theme_version"`
}
