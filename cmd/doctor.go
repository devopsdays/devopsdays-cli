// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

// doctorCmd represents the doctor command
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check that everything looks good",
	Long:  `Use the doctor command to evaluate your environment and make sure that everything is ready for you to start doing some stuff.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("Checking your config...")
		out, err := exec.Command("hugo", "version").Output()
		if err != nil {
			log.Fatal(err)
		}
		s := string(out[:])
		re := regexp.MustCompile("v....")
		hugoVersion := re.FindString(s)
		switch hugoVersion {
		case "v0.16":
			fmt.Println("\u2713 Hugo version", hugoVersion, "is okay")
		case "v0.17":
			fmt.Println("\u2713 Hugo version", hugoVersion, "is okay")
		default:
			fmt.Println("\u2717 Hugo version", hugoVersion, "is incompatible. Please use a supported version (0.16 or 0.17)")
		}
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
