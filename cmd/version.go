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
	"path/filepath"
	"regexp"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

var (
	Version = "master"
	Build   string
)

func init() {
	RootCmd.AddCommand(versionCmd)
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

type Theme struct {
	Version string `toml:"theme_version"`
}
