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
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

// eventCmd represents the event command
var eventCmd = &cobra.Command{
	Use:   "event [city year]",
	Short: "Create a new event",
	Long: `Create a new event.
The 'city' and 'year' arguments are optional, but if you provide year, you must also provide city.
City must not have a space. Replace spaces with '-'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			addEvent(args[0])
		} else {
			addEvent("")
		}
	},
}

func init() {
	RootCmd.AddCommand(eventCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// eventCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// eventCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// addEvent creates a new event based upon city, year, and twitter handle.
// It returns an empty string and an error if the event already exists.
func addEvent(city string) (err error) {

	reader := bufio.NewReader(os.Stdin) // TODO: Convert to a loop for each argument - maybe a map?
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
	if validateField(strings.TrimSpace(eventYear), "year") == false {
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
	if result, err := createEventFile(city, eventYear); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Event created for %s!!!\n", result)
	}
	return
}

func createEventFile(city, year string) (string, error) {

	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)}
	slug := strings.Join(s, "")
	// cityClean := strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)
	t := template.Must(template.New("event.yml.tmpl").ParseFiles("templates/event.yml.tmpl"))
	data := struct {
		City      string
		Year      string
		Slug      string
		CityClean string
	}{
		city,
		strings.TrimSpace(year),
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
		}
		return true
	case "year":
		if strings.Count(input, "") != 5 {
			return false
		} else if s, err := strconv.ParseInt(input, 10, 32); err == nil {
			if s < 2016 || s > 2030 {
				return false
			}
			return true

		}
	case "twitter":
		if strings.ContainsAny(input, " ") {
			return false
		}
		return true
	}
	return true // TODO: Make this return an error if no field was matched
}
