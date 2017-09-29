// The MIT License (MIT)
// Copyright (c) 2017 Matt Stratton
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

// const supportedVersion = "v0.26"

// doctorCmd represents the doctor command
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check that everything looks good",
	Long:  `Use the doctor command to evaluate your environment and make sure that everything is ready for you to start doing some stuff.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("Checking your config...")
		checkHugo()
		checkGit()
	},
}

func init() {
	RootCmd.AddCommand(doctorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doctorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doctorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func checkHugo() {
	supportedVersions := map[string]bool{"0.23": true, "0.24.1": true, "0.25.1": true, "0.26": true, "0.27": true, "0.28": true, "0.29": true}
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
		for k, _ := range supportedVersions {
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
