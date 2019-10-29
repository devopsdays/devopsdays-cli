package commands

// cSpell:disable
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

// cSpell:enable

var (
	// Version is the current version of the devopsdays-cli tool. Unless set elsewhere, it is referred to as "master"
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
	},
}

func showVersion() {
	fmt.Println("devopsdays-cli version: ", Version)
	fmt.Println("devopsdays-cli build: ", Build)
	fmt.Println("hugo version: ", getHugoVersion())
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func getHugoVersion() (hugoVersion string) {
	out, err := exec.Command("hugo", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out[:])
	re := regexp.MustCompile("([0-9]\\.)\\w+")
	hugoVersion = re.FindString(s)
	return
}

func getCliVersion() (cliVersion string) {
	return Version
}

func chompVersion(version string) string {
	return regexp.MustCompile("([0-9]\\.)\\w+").FindString(version)
}

// Theme represents the currently installed devopsdays-theme Hugo theme.
// The field Version represents the current version.
type Theme struct {
	Version string `toml:"theme_version"`
}
