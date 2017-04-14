// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

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
		fmt.Println("editEvent called")
		// TODO: Check for args first
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the city:")
		city, _ := reader.ReadString('\n')
		fmt.Println("Enter the year:")
		year, _ := reader.ReadString('\n')
		if checkEvent(city, year) == false {
			fmt.Println("That city does not exist.")
		}
		myEvent := eventStruct(city, year)
		editEvent(myEvent)
		// fmt.Println(fieldMap())
		// spew.Dump(myEvent)
	},
}

func init() {
	eventCmd.AddCommand(editEventCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editEventCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editEventCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

type Event struct {
	Name                  string `yaml:"name"`
	Year                  string `yaml:"year"`
	City                  string `yaml:"city"`
	EventTwitter          string `yaml:"event_twitter"`
	Description           string `yaml:"description"`
	GaTrackingID          string `yaml:"ga_tracking_id"`
	Startdate             string `yaml:"startdate"`
	Enddate               string `yaml:"enddate"`
	CfpDateStart          string `yaml:"cfp_date_start"`
	CfpDateEnd            string `yaml:"cfp_date_end"`
	CfpDateAnnounce       string `yaml:"cfp_date_announce"`
	CfpOpen               string `yaml:"cfp_open"`
	CfpLink               string `yaml:"cfp_link"`
	RegistrationDateStart string `yaml:"registration_date_start"`
	RegistrationDateEnd   string `yaml:"registration_date_end"`
	RegistrationClosed    string `yaml:"registration_closed"`
	RegistrationLink      string `yaml:"registration_link"`
	Coordinates           string `yaml:"coordinates"`
	Location              string `yaml:"location"`
	LocationAddress       string `yaml:"location_address"`
	NavElements           []struct {
		Name string `yaml:"name"`
	} `yaml:"nav_elements"`
	TeamMembers []struct {
		Name     string `yaml:"name"`
		Twitter  string `yaml:"twitter,omitempty"`
		Employer string `yaml:"employer,omitempty"`
		Github   string `yaml:"github,omitempty"`
		Facebook string `yaml:"facebook,omitempty"`
		Linkedin string `yaml:"linkedin,omitempty"`
		Website  string `yaml:"website,omitempty"`
		Image    string `yaml:"image,omitempty"`
	} `yaml:"team_members"`
	OrganizerEmail string `yaml:"organizer_email"`
	ProposalEmail  string `yaml:"proposal_email"`
	Sponsors       []struct {
		ID    string `yaml:"id"`
		Level string `yaml:"level"`
	} `yaml:"sponsors"`
	SponsorsAccepted string `yaml:"sponsors_accepted"`
	SponsorLevels    []struct {
		ID    string `yaml:"id"`
		Label string `yaml:"label"`
		Max   int    `yaml:"max,omitempty"`
	} `yaml:"sponsor_levels"`
}

func fieldMap() (fieldMap map[string]string) {
	tempMap := make(map[string]string)
	tempMap["EventTwitter"] = "Twitter"
	tempMap["GaTrackingID"] = "Google Analytics Tracking ID"
	tempMap["Startdate"] = "Start Date"
	tempMap["Enddate"] = "End Date"
	tempMap["CfpDateStart"] = "CFP Start Date"
	tempMap["CfpDateEnd"] = "CFP End Date"
	tempMap["CfpDateAnnounce"] = "CFP Announcement Date"
	tempMap["CfpOpen"] = "CFP Link"
	tempMap["RegistrationDateStart"] = "Registation Start Date"
	tempMap["RegistrationDateEnd"] = "Registration End Date"
	tempMap["RegistrationLink"] = "Registration Link"
	tempMap["Coordinates"] = "Coordinates"
	tempMap["Location"] = "Location"
	tempMap["LocationAddress"] = "Location Address"

	return tempMap
}

func eventFields() []string {
	fields := make([]string, 14)
	fields[0] = "EventTwitter"
	fields[1] = "GaTrackingID"
	fields[2] = "Startdate"
	fields[3] = "Enddate"
	fields[4] = "CfpDateStart"
	fields[5] = "CfpDateEnd"
	fields[6] = "CfpDateAnnounce"
	fields[7] = "CfpOpen"
	fields[8] = "RegistrationDateStart"
	fields[9] = "RegistrationDateEnd"
	fields[10] = "RegistrationLink"
	fields[11] = "Coordinates"
	fields[12] = "Location"
	fields[13] = "LocationAddress"

	return fields
}

func editEvent(event Event) (err error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you want to [1] edit the value of a field, [2] add an organizer, or [3] add a sponsor?")
	c, _ := reader.ReadString('\n')
	c = strings.TrimSpace(c)
	switch c {
	case "1":
		s := eventFields()
		myField := makeMenu(s)
		fmt.Println("The value of this field is: ", returnField(event, myField))
		fmt.Println("Would you like to change it?")
		c, _ := reader.ReadString('\n')
		c = strings.TrimSpace(c)
		if c == "y" {
			fmt.Println("What would you like to change it to?")
			c, _ := reader.ReadString('\n')
			c = strings.TrimSpace(c)
			editField(event, myField, c)
		}
	case "2":
		fmt.Println("Adding organizers is not yet supported.")
	case "3":
		fmt.Println("Adding sponsors is not yet supported.")
	default:
		fmt.Println("This is the default.")
	}

	// Note: This is commented out for now, but we want to use this functionality somewhere
	// event.Name = "mugsyville"
	// fmt.Print(event.Name)
	// y, err := yaml.Marshal(&event)
	// ioutil.WriteFile((eventDataPath(webdir, city, year)), y, 0755)

	return

}

func eventStruct(city, year string) (event Event) {
	// var event Event
	yamlFile, err := ioutil.ReadFile(eventDataPath(webdir, city, year))
	err = yaml.Unmarshal(yamlFile, &event)
	if err != nil {
		panic(err)
	}
	return event
}

// TODO: This should actually return the key to change; rather than just create the menu
func makeMenu(items []string) (field string) {
	fmt.Println("Which field would you like to modify?")
	myMap := fieldMap()
	menu := "\n"
	for i, v := range items {
		menu += "["
		menu += strconv.Itoa(i + 1)
		menu += "] "
		menu += myMap[v]
		menu += "\n"
	}
	fmt.Println(menu)
	reader := bufio.NewReader(os.Stdin)
	var c, _ = reader.ReadString('\n')
	c = strings.TrimSpace(c)
	c2, _ := strconv.Atoi(c)
	field = items[(c2 - 1)]

	return field
}

func returnField(event Event, field string) (name string) {
	r := reflect.ValueOf(event)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func editField(event Event, field, value string) {
	// r := reflect.ValueOf(event)
	// f := reflect.Indirect(r).FieldByName(field)
	reflect.ValueOf(&event).Elem().FieldByName(field).SetString(value)
	y, _ := yaml.Marshal(&event)
	ioutil.WriteFile((eventDataPath(webdir, event.City, event.Year)), y, 0755)
	return
}
