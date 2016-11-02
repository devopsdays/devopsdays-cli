// Copyright 2016 Matt Stratton. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main sets up the probablyfine application.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/nfnt/resize"
	"gopkg.in/urfave/cli.v1"
)

const version = "0.1.1"
const webdir = "/Users/mattstratton/src/devopsdays-web" // TODO: Change this to read an environment variable, and default to cwd if envar not set

func main() {
	app := cli.NewApp()
	app.Version = version
	app.Name = "probablyfine"
	app.Usage = "Run maintainence tasks for the devopsdays.org website"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Matt Stratton",
			Email: "matt.stratton@gmail.com",
		},
	}
	app.Copyright = "(c) 2016 Matt Stratton"
	app.HelpName = "probablyfine"
	app.Commands = []cli.Command{
		{
			Name:    "event",
			Aliases: []string{"e"},
			Usage:   "options for events",
			Subcommands: []cli.Command{
				{
					Name:        "add",
					Usage:       "add a new event",
					Description: "Adds a new event. Takes the city name as an argument. Put the city name in quotes if there are spaces",
					ArgsUsage:   "[cityname, year]",
					Action: func(c *cli.Context) error {
						city := c.Args().Get(0) // TODO: Add ability to take year and twitter as arguments
						// fmt.Printf("new event for %s added\n", city)
						if err := addEvent(city); err != nil {
							fmt.Printf("Error: %s\n", err)
						} else {
							fmt.Println("New event for", strings.TrimSpace(city)) // TODO: Add full values being returned
						}
						return nil
					},
				},
			},
		},
		{
			Name:    "sponsor",
			Aliases: []string{"s"},
			Usage:   "options for sponsors",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new sponsor",
					Action: func(c *cli.Context) error {
						sponsor := c.Args().Get(0)
						if sponsor == "" {
							fmt.Println("You must specify a sponsor name") //TODO: Make this an error
							return nil
						}
						if c.Args().Get(1) != "" {
							fmt.Println("Sponsors must not have spaces")
							return nil
						}
						if err := addSponsor(sponsor); err != nil {
							fmt.Printf("Error: %s\n", err)
						} else {
							fmt.Println("New sponsor for", sponsor, "created.") // TODO: Add full values being returned
						}
						return nil
					},
				},
				{
					Name:  "audit",
					Usage: "audit all sponsors for logos and proper size",
					Action: func(c *cli.Context) error {
						logmsg := "All sponsors look fine"
						fmt.Printf("%s\n", logmsg)
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

// addEvent creates a new event based upon city, year, and twitter handle.
// It returns an empty string and an error if the event already exists.
func addEvent(city string) (err error) {

	reader := bufio.NewReader(os.Stdin) // TODO: Convert to a loop for each argument
	if city == "" {
		fmt.Println("Enter the city: ")
		city, _ = reader.ReadString('\n')
	}
	if validateField(city, "city") == false {
		return fmt.Errorf("That is an invalid city. It should be less than 100 characters.")
	}
	t := time.Now()
	fmt.Printf("Enter your event year (default %s): ", t.Format("2006")) // TODO: Prompt user to keep trying on invalid entry
	eventYear, _ := reader.ReadString('\n')
	if eventYear == "\n" {
		eventYear = t.Format("2006")
	}
	if validateField(eventYear, "year") == false {
		return fmt.Errorf("That is an invalid year. It must be four digits and between 2016 and 2030.")
	}
	fmt.Println("Enter your devopsdays event twitter handle (defaults to devopsdays): ")
	eventTwitter, _ := reader.ReadString('\n')
	if eventTwitter == "\n" {
		eventTwitter = "devopsdays"
	} else {
		eventTwitter = strings.TrimSpace(strings.Replace(eventTwitter, "@", "", 1))
	}
	if validateField(eventTwitter, "twitter") == false {
		return fmt.Errorf("That is an invalid Twitter handle. It must not contain spaces.")
	}

	// build the event data file path
	s := []string{webdir, "/data/events/", strings.TrimSpace(eventYear), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10), ".yml"}
	eventDataPath := strings.Join(s, "")
	if _, err := os.Stat(eventDataPath); err == nil {
		return fmt.Errorf("The event already exists")
	}

	// create the event file
	if result, err := createEventFile(city, eventYear, eventTwitter); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Event created for %s!!!", result)
	}
	return
}

func createEventFile(city, year, twitter string) (string, error) {

	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)}
	slug := strings.Join(s, "")
	// cityClean := strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)
	t := template.Must(template.New("event.yml.tmpl").ParseFiles("templates/event.yml.tmpl"))
	data := struct {
		City      string
		Year      string
		Twitter   string
		Slug      string
		CityClean string
	}{
		city,
		strings.TrimSpace(year),
		twitter,
		slug,
		cityClean(city),
	}
	f, err := os.Create(eventDataPath(city, year))
	if err != nil {
		return "", err
	}
	defer f.Close()
	t.Execute(f, data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created event file for", city, "for year", year, "at", eventDataPath(city, year))
	}
	return city, nil
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

func addSponsor(sponsor string) (err error) { // TODO: write addSponsor() function

	// Check if the sponsor exists already (use the checkSponsor() function) TODO: Write checkSponsor() function
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
	// check if the sponsor image file meets requirements using checkSponsorImageSize() TODO: write checkSponsorImageSize() function

	// if sponsor image doesn't meet requirements, offer to resize it using resizeImage() TODO: write resizeImage()
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

func cityClean(city string) (cityClean string) {
	cityClean = strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)
	return
}

func eventDataPath(city, year string) (eventDataPath string) { // TODO: Add argument for webdir path
	s := []string{webdir, "/data/events/", strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10), ".yml"}
	eventDataPath = strings.Join(s, "")
	return eventDataPath
}

func validateField(input, field string) bool { // TODO: Write validateField() function
	switch field {
	case "city":
		if strings.Count(input, "") > 100 {
			return false
		} else {
			return true
		}
	case "year":
		if strings.Count(input, "") != 5 {
			return false
		} else if s, err := strconv.ParseInt(input, 10, 32); err == nil {
			if s < 2016 || s > 2030 {
				return false
			} else {
				return true
			}
		}
	case "twitter":
		if strings.ContainsAny(input, " ") {
			return false
		}
		return true
	}
	return true // TODO: Make this return an error if no field was matched
}

// checkSponsor takes in one argument, the name of a sponsor, and returns true if the sponsor already exists.
func checkSponsor(sponsor string) bool {
	fmt.Println(sponsorDataPath(webdir, sponsor))
	if _, err := os.Stat(sponsorDataPath(webdir, sponsor)); err == nil {
		return true
	} else {
		return false
	}
}

func checkSponsorImage(path string) bool {
	fmt.Println(path)
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
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
