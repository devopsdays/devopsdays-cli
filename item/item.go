// Package item contains the types and methods for items like talks, events, sponsors, etc.
package item

import "fmt"

// Item is the overall interface for the items (wow, that's terrible)
type Item interface {
	Create(city, year string) (*struct{}, error)
}

// Event is an event
type Event struct {
	Name                  string   `yaml:"name"`
	Year                  string   `yaml:"year"`
	City                  string   `yaml:"city"`
	EventTwitter          string   `yaml:"event_twitter"`
	Description           string   `yaml:"description"`
	GoogleAnalytics       string   `yaml:"ga_tracking_id"`
	StartDate             string   `yaml:"startdate"`
	EndDate               string   `yaml:"enddate"`
	CFPDateStart          string   `yaml:"cfp_date_start"`
	CFPDateEnd            string   `yaml:"cfp_date_end"`
	CFPDateAnnounce       string   `yaml:"cfp_date_announce"`
	CFPOpen               string   `yaml:"cfp_open"`
	CFPLink               string   `yaml:"cfp_link"`
	RegistrationDateStart string   `yaml:"registration_date_start"`
	RegistrationDateEnd   string   `yaml:"registration_date_end"`
	RegistrationClosed    string   `yaml:"registration_closed"`
	RegistrationLink      string   `yaml:"registration_link"`
	MastheadBackground    string   `yaml:"masthead_background"`
	Coordinates           string   `yaml:"coordinates"`
	Location              string   `yaml:"location"`
	LocationAddress       string   `yaml:"location_address"`
	NavElements           []string `yaml:"nav_elements"`
	TeamMembers           []string `yaml:"team_members"`
	OrganizerEmail        string   `yaml:"organizer_email"`
	ProposalEmail         string   `yaml:"proposal_email"`
	Sponsors              []string `yaml:"sponsors"`
	SponsorsAccepted      string   `yaml:"sponsors_accepted"`
	SponsorLevels         []string `yaml:"sponsor_levels"`
}

func (e *Event) Create(city, year string) (*Event, error) {
	myEvent := new(Event)
	myEvent.Name = "Hello"
	return myEvent, nil
}

func ShowEvent() {
	e := new(Event)
	event, _ := e.Create("Ponyville", "2017")
	fmt.Println(event)
}
