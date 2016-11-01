package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"gopkg.in/urfave/cli.v1"
)

const version = "0.1.1"
const webdir = "/Users/mattstratton/src/devopsdays-web"

func main() {
	app := cli.NewApp()
	app.Version = version
	app.Name = "Probably Fine"
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
						sponsor := "Chef"
						fmt.Printf("new sponsor for %s added\n", sponsor)
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
	t := time.Now()
	fmt.Printf("Enter your event year (default %s): ", t.Format("2006")) // TODO: Prompt user to keep trying on invalid entry
	eventYear, _ := reader.ReadString('\n')
	if eventYear == "\n" {
		eventYear = t.Format("2006")
	}
	if validateField(eventYear, "year") == false {
		return fmt.Errorf("That is an invalid year")
	}
	fmt.Println("Enter your devopsdays event twitter handle (defaults to devopsdays): ")
	eventTwitter, _ := reader.ReadString('\n')
	if eventTwitter == "\n" {
		eventTwitter = "devopsdays"
	} else {
		eventTwitter = strings.TrimSpace(strings.Replace(eventTwitter, "@", "", 1))
	}
	// Just some output for now
	// fmt.Println("New event for", strings.TrimSpace(city), "in", strings.TrimSpace(eventYear))
	// build the event data file path
	s := []string{webdir, "/data/events/", strings.TrimSpace(eventYear), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10), ".yml"}
	eventDataPath := strings.Join(s, "")
	if _, err := os.Stat(eventDataPath); err == nil {
		return fmt.Errorf("The event already exists")
	}
	// fmt.Println("The friendly name is ", strings.Replace(city, " ", "-", 2))
	// fmt.Println("The data path is ", eventDataPath)
	// fmt.Println("The twitter ID is ", eventTwitter)

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
	t := template.Must(template.New("event.yml.tmpl").ParseFiles("event.yml.tmpl"))
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

func addSponsor(sponsor string) { // TODO: write addSponsor() function
	// Read in the sponsor name

	// Check if the sponsor exists already (use the checkSponsor() function) TODO: Write checkSponsor() function

	// If the sponsor already exists, prompt the user for a new name, suggesting to append "-YYYY" after the sponsors name

	// check if the new sponsor suggested name exists

	// prompt for the path to the sponsor image file

	// check if the sponsor image file meets requirements using checkSponsorImageSize() TODO: write checkSponsorImageSize() function

	// if sponsor image doesn't meet requirements, offer to resize it using resizeImage() TODO: write resizeImage()

	// prompt for sponsor URL

	// write sponsor YAML file and copy image from path to proper destination
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
			if s < 2016 {
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
	return true
}
