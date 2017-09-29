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
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/spf13/cobra"
)

// eventCmd represents the "event" command
var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// TODO: Work your own magic here
	// 	fmt.Println("speaker called")
	// },
}

// createEventCmd represents the "event create" command
var createEventCmd = &cobra.Command{
	Use:   "event create [city year]",
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

// editEventCmd represents the editEvent command
var editEventCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		// TODO: Check for args first
		city := cityFlag
		year := yearFlag
		if city != "" {
			if checkEvent(city, year) == false {
				log.Fatal("That city does not exist.")
			}
			myEvent := eventStruct(city, year)
			editEvent(myEvent)
		} else {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Enter the city:")
			city, _ := reader.ReadString('\n')
			fmt.Println("Enter the year:")
			year, _ := reader.ReadString('\n')
			if checkEvent(city, year) == false {
				log.Fatal("That city does not exist.")
			}
			myEvent := eventStruct(city, year)
			editEvent(myEvent)
		}

	},
}

func init() {
	RootCmd.AddCommand(eventCmd)
	eventCmd.AddCommand(createEventCmd)
	eventCmd.AddCommand(editEventCmd)

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
	if result, err := createEventFile(city, eventYear, eventTwitter); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Event created for %s!!!\n", result)
	}

	// create the event content directory
	if result, err := createEventContentDir(city, eventYear); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Event content directory created for %s!!!\n", result)
	}

	// create the event content files
	contentfiles := []string{"index", "conduct", "contact", "location", "program", "propose", "registration", "sponsor"}
	for _, contentFile := range contentfiles {

		if result, err := createEventContentFile(city, eventYear, contentFile); err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("Event content file created for %s!!!\n", result)
		}

	}

	return
}

func createEventFile(city, year, twitter string) (string, error) {

	// find a rice.Box
	templateBox, err := rice.FindBox("../templates")
	if err != nil {
		log.Fatal(err)
	}
	// get file contents as string
	templateString, err := templateBox.String("event.yml.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)}
	slug := strings.Join(s, "")
	// cityClean := strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)
	// t := template.Must(template.New("event.yml.tmpl").ParseFile("templates/event.yml.tmpl"))
	// parse and execute the template
	t, err := template.New("event.yml").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		City      string
		Year      string
		Slug      string
		CityClean string
		Twitter   string
	}{
		strings.TrimSpace(city),
		strings.TrimSpace(year),
		slug,
		cityClean(city),
		strings.TrimSpace(twitter),
	}
	f, err := os.Create(eventDataPath(webdir, city, year))
	if err != nil {
		return "", err
	}
	defer f.Close()
	t.Execute(f, data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created event file for", city, "for year", year, "at", eventDataPath(webdir, city, year))
	}
	return city, nil
}

func createEventContentDir(city, year string) (string, error) {
	err := os.MkdirAll((eventContentPath(webdir, city, year)), 0755)
	if err != nil {
		return "", err
	}
	return city, nil
}

func createEventContentFile(city, year, page string) (string, error) { // add page as an argument later

	// find a rice.Box
	templateBox, err := rice.FindBox("../templates")
	if err != nil {
		log.Fatal(err)
	}
	templateName := "events/" + page + ".md.tmpl"
	// templateName := "index.md.tmpl"
	// get file contents as string
	templateString, err := templateBox.String(templateName)
	if err != nil {
		// log.Fatal(templateName)
		log.Fatal(err)
	}
	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)}
	slug := strings.Join(s, "")
	// t := template.Must(template.New(page+".md.tmpl").Delims("[[", "]]").ParseFiles(templateString))
	// parse and execute the template
	t, err := template.New(page+".md").Delims("[[", "]]").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		City      string
		Year      string
		Slug      string
		CityClean string
	}{
		strings.TrimSpace(city),
		strings.TrimSpace(year),
		slug,
		cityClean(city),
	}
	filePath := filepath.Join((eventContentPath(webdir, city, year)), (page + ".md"))
	f, err := os.Create(filePath)
	if err != nil {
		return "Cannot create", err
	}
	defer f.Close()
	t.Execute(f, data)
	if err != nil {
		fmt.Println(err, "template execute error")
	} else {
		fmt.Println("Created event content file for", city, "for year", year, "at", filePath)
	}
	return city, nil

}
