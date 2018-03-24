package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Returns the current configuration",
	Long:  `Displays any environment variables and configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("Current configuration")
		fmt.Println("DODPATH = ", os.Getenv("DODPATH"))
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Current Working Directory = ", pwd)
		fmt.Println("DevOpsDays web directory = ", webdir)
		color.Blue("Checking your config...")
		CheckHugo()
		checkGit()
	},
}

func init() {
	showCmd.AddCommand(configCmd)

}

// CheckHugo tests whether or not a compatible version of the Hugo static site generator is instealled.
//
// Currently, the list of supported versions is hard-coded using the `supportedVersions` variable, but this should be moved elsewhere eventually.
func CheckHugo() {
	supportedVersions := map[string]bool{"0.36.1": true, "0.37": true, "0.37.1": true}
	out, err := exec.Command("hugo", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out[:])
	re := regexp.MustCompile(`[0-9]+(\.[0-9]+)*`)
	hugoVersion := re.FindString(s)
	if supportedVersions[hugoVersion] {
		fmt.Println("\u2713 Hugo version", hugoVersion, "is okay")
	} else {
		fmt.Println("\u2717 Hugo version", hugoVersion, "is incompatible.")
		fmt.Println("Supported Versions are:")
		for k := range supportedVersions {
			fmt.Println(k)
		}
	}
}

func checkGit() {
	_, err := exec.Command("git", "version").Output()
	if err != nil {
		fmt.Println("\u2717 git is not installed")
	} else {
		fmt.Println("\u2713 git is installed")
	}
}
