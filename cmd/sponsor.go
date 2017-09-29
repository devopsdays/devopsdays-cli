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
	"bufio"
	"errors"
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/nfnt/resize"
	"github.com/spf13/cobra"
)

// sponsorCmd represents the sponsor command
var sponsorCmd = &cobra.Command{
	Use:   "sponsor [name]",
	Short: "Create a sponsor",
	Long: `Create a new sponsor file, and optionally add the sponsor's image.
The name argument must not contain any spaces.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			addSponsor(args[0])
		} else {
			addSponsor("")
		}
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

func addSponsor(sponsor string) (err error) {

	reader := bufio.NewReader(os.Stdin)
	if sponsor == "" {
		fmt.Println("Enter the sponsor's name. It must not contain any spaces: ")
		sponsor, _ = reader.ReadString('\n')
	}
	// Check if the sponsor exists already
	if checkSponsor(sponsor) == true {
		return errors.New("Sponsor already exists. Try adding it again, perhaps appending '-YYYY'\nFor example, 'chef-2017'")
	}
	// prompt for the path to the sponsor image file
	fmt.Println("Optional: Enter the path to the sponsor's image. It must be the full path. For example: `/Users/mattstratton/chef.png`. Enter return to add the sponsor image manually later.")
	sponsorImage, _ := reader.ReadString('\n')
	if sponsorImage == "\n" {
		fmt.Println("No sponsor image entered. Be sure to copy it to the path ", sponsorImagePath(webdir, sponsor), "later.")
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

func createSponsorFile(sponsor, sponsorName, sponsorUrl string) (string, error) {
	// find a rice.Box
	templateBox, err := rice.FindBox("../templates")
	if err != nil {
		log.Fatal(err)
	}
	// get file contents as string
	templateString, err := templateBox.String("sponsor.yml.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	// t := template.Must(template.New("sponsor.yml.tmpl").ParseFiles("templates/sponsor.yml.tmpl"))
	t, err := template.New("sponsor.yml").Parse(templateString)
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
	s := []string{strings.TrimSpace(sponsor), ".yml"}
	// sponsorDataPath = strings.Join(s, "")
	sponsorDataPath = filepath.Join(webdir, "data", "sponsors", strings.Join(s, ""))
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

	m := resize.Resize(600, 0, img, resize.Lanczos3)

	out, err := os.Create(destPath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	png.Encode(out, m)
}
