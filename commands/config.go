package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"

	"github.com/spf13/cobra"

	// "github.com/pkg/errors"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
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
		checkHugo()
		checkGit()
	},
}

func init() {
	showCmd.AddCommand(configCmd)

}

func showConfig() {
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
	checkHugo()
	checkGit()
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// checkHugo tests whether or not a compatible version of the Hugo static site generator is instealled.
//
// Currently, the list of supported versions is hard-coded using the `supportedVersions` variable, but this should be moved elsewhere eventually.
func checkHugo() {
	// supportedVersions := map[string]bool{"0.36.1": true, "0.37": true, "0.37.1": true}
	currentThemeVersion, currentHugoVersion, currentCliVersion, err := getCurrentVersions()
	if err != nil {
		log.Fatal(err)
	}

	if chompVersion(currentHugoVersion) == chompVersion(getHugoVersion()) {
		color.Green("\u2713 Hugo version %s is okay", getHugoVersion())
	} else {
		color.Red("\u2717 Hugo version %s is incompatible.", getHugoVersion())
		color.Red("Supported Version is: %s", currentHugoVersion)
	}

	if chompVersion(currentThemeVersion) == chompVersion(getThemeVersion()) {
		color.Green("\u2713 Theme version %s is okay", getThemeVersion())
	} else {
		color.Red("\u2717 Theme version %s is incompatible.", getThemeVersion())
		color.Red("Supported Version is: %s", currentThemeVersion)
	}

	if currentCliVersion == Version {
		color.Green("\u2713 devopsdays-cli version %s is okay", Version)
	} else {
		color.Red("\u2717 devopsdays-cli version %s is incompatible.", Version)
		color.Red("Supported Version is: %s", currentCliVersion)
	}
}

type currentVersion struct {
	ThemeVersion         string `json:"theme_version"`
	HugoVersion          string `json:"hugo_version"`
	DevopsdaysCliVersion string `json:"devopsdays_cli_version"`
}

// getCurrentVersions returns the supported version of the devopsdays Hugo theme, hugo, and devopsdays-cli
func getCurrentVersions() (themeVersion string, hugoVersion string, devopsdaysCliVersion string, err error) {

	// url := "https://rawgit.com/devopsdays/devopsdays-web/master/metadata.json"
	url := "https://raw.githubusercontent.com/devopsdays/devopsdays-web/refs/heads/main/metadata.json"

	devopsdaysClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "devopsdays-cli")

	res, getErr := devopsdaysClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	myCurrentVersion := currentVersion{}
	jsonErr := json.Unmarshal(body, &myCurrentVersion)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return myCurrentVersion.ThemeVersion, myCurrentVersion.HugoVersion, myCurrentVersion.DevopsdaysCliVersion, nil
}

func checkGit() {
	_, err := exec.Command("git", "version").Output()
	if err != nil {
		fmt.Println("\u2717 git is not installed")
	} else {
		fmt.Println("\u2713 git is installed")
	}
}

func compareVersion(currentVersion string, localVersion string) bool {
	if chompVersion(currentVersion) == chompVersion(localVersion) {
		return true
	}
	return false
}
