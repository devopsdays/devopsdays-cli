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
	"fmt"
	"io/ioutil"

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
		editEvent("ponyville", "2018")
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
	Name                  string      `yaml:"name"`
	Year                  string      `yaml:"year"`
	City                  string      `yaml:"city"`
	EventTwitter          string      `yaml:"event_twitter"`
	Description           string      `yaml:"description"`
	GaTrackingID          string      `yaml:"ga_tracking_id"`
	Startdate             interface{} `yaml:"startdate"`
	Enddate               interface{} `yaml:"enddate"`
	CfpDateStart          interface{} `yaml:"cfp_date_start"`
	CfpDateEnd            interface{} `yaml:"cfp_date_end"`
	CfpDateAnnounce       interface{} `yaml:"cfp_date_announce"`
	CfpOpen               string      `yaml:"cfp_open"`
	CfpLink               string      `yaml:"cfp_link"`
	RegistrationDateStart interface{} `yaml:"registration_date_start"`
	RegistrationDateEnd   interface{} `yaml:"registration_date_end"`
	RegistrationClosed    string      `yaml:"registration_closed"`
	RegistrationLink      string      `yaml:"registration_link"`
	Coordinates           string      `yaml:"coordinates"`
	Location              string      `yaml:"location"`
	LocationAddress       string      `yaml:"location_address"`
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

func editEvent(city, year string) (err error) {
	var event Event
	yamlFile, err := ioutil.ReadFile(eventDataPath(webdir, city, year))
	err = yaml.Unmarshal(yamlFile, &event)
	if err != nil {
		panic(err)
	}
	event.Name = "mugsyville"
	fmt.Print(event.Name)
	y, err := yaml.Marshal(&event)
	ioutil.WriteFile((eventDataPath(webdir, city, year)), y, 0755)

	return

}

// func eventDataPath(webdir, city, year string) (eventDataPath string) { // TODO: Add argument for webdir path
// 	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10), ".yml"}
// 	eventDataPath = filepath.Join(webdir, "data", "events", strings.Join(s, ""))
// 	// eventDataPath = strings.Join(s, "")
// 	// eventDataPath = webdir
// 	return eventDataPath
// }
