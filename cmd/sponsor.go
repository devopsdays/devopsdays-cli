// Copyright © 2016 Matt Stratton <matt.stratton@gmail.com>
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
	"bufio"
	"errors"
	"fmt"
	"image/png"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/nfnt/resize"
	"github.com/spf13/cobra"
)

// sponsorCmd represents the sponsor command
var sponsorCmd = &cobra.Command{
	Use:   "sponsor [name]",
	Short: "Create a new sponsor",
	Long: `Create a new sponsor file, and optionally add the sponsor's image.
The name argument must not contain any spaces.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("sponsor called")
	},
}

func init() {
	RootCmd.AddCommand(sponsorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sponsorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sponsorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func createSponsorFile(sponsor, sponsorName, sponsorUrl string) (string, error) {
	t := template.Must(template.New("sponsor.yml.tmpl").ParseFiles("templates/sponsor.yml.tmpl"))
	data := struct {
		Name string
		Url  string
	}{
		strings.TrimSpace(sponsorName),
		strings.TrimSpace(sponsorUrl),
	}
	f, err := os.Create(sponsorDataPath(webdir, sponsor))
	if err != nil {
		return "", err
	}
	defer f.Close()
	t.Execute(f, data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created sponsor file for", sponsor, "at", sponsorDataPath(webdir, sponsor))
	}
	return sponsor, nil

}

func addSponsor(sponsor string) (err error) {

	// Check if the sponsor exists already
	if checkSponsor(sponsor) == true {
		return errors.New("Sponsor already exists. Try adding it again, perhaps appending '-YYYY'\nFor example, 'chef-2017'")
	}

	reader := bufio.NewReader(os.Stdin)
	// prompt for the path to the sponsor image file
	fmt.Println("Optional: Enter the path to the sponsor's image. It must be the full path. For example: `/Users/mattstratton/chef.png`. Enter return to add the sponsor image manually later.")
	sponsorImage, _ := reader.ReadString('\n')
	if sponsorImage == "\n" {
		fmt.Println("No sponsor image found. Be sure to copy it to the path ", sponsorImagePath(webdir, sponsor), "later.")
	} else {

		if sponsorImage = strings.TrimSpace(sponsorImage); checkSponsorImage(sponsorImage) == false {
			return errors.New("Sponsor image not found.")
		}
	}

	// prompt for sponsor's name
	fmt.Println("Enter the sponsor's full name. For example: `Chef Software, Inc`")
	sponsorName, _ := reader.ReadString('\n')
	if sponsorName == "\n" {
		return errors.New("Sponsor Name is required.")
	}
	fmt.Println(sponsorName)

	// prompt for sponsor URL
	fmt.Println("Enter the sponsor's URL. It must include 'http://' or 'https://'. For example: `https://www.chef.io`")
	sponsorUrl, _ := reader.ReadString('\n')
	if sponsorUrl == "\n" {
		return errors.New("Sponsor URL is required.")
	}

	// write sponsor YAML file and copy image from path to proper destination
	createSponsorFile(sponsor, sponsorName, sponsorUrl)
	fmt.Println("Sponsor created for ", sponsorName)
	if sponsorImage != "\n" {
		resizeSponsorImage(strings.TrimSpace(sponsorImage), sponsorImagePath(webdir, sponsor))
	} else {
		fmt.Println("Don't forget to place the sponsor image at ", sponsorImagePath(webdir, sponsor))
	}
	return
}

// checkSponsor takes in one argument, the name of a sponsor, and returns true if the sponsor already exists.
func checkSponsor(sponsor string) bool {
	fmt.Println(sponsorDataPath(webdir, sponsor))
	if _, err := os.Stat(sponsorDataPath(webdir, sponsor)); err == nil {
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

func sponsorDataPath(webdir, sponsor string) (sponsorDataPath string) {
	s := []string{webdir, "/data/sponsors/", strings.TrimSpace(sponsor), ".yml"}
	sponsorDataPath = strings.Join(s, "")
	return sponsorDataPath
}

func sponsorImagePath(webdir, sponsor string) (sponsorImagePath string) {
	s := []string{webdir, "/static/img/sponsors/", strings.TrimSpace(sponsor), ".png"}
	sponsorImagePath = strings.Join(s, "")
	return sponsorImagePath
}

func resizeSponsorImage(srcPath, destPath string) {
	fmt.Println("Resizing image")
	file, err := os.Open(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(200, 200, img, resize.Lanczos3)

	out, err := os.Create(destPath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	png.Encode(out, m)
}
