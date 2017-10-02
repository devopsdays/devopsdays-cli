package model

// Event defines a devopsdays event in the data yaml file
type Event struct {
	Name                  string         `yaml:"name"`
	Year                  string         `yaml:"year"`
	City                  string         `yaml:"city"`
	EventTwitter          string         `yaml:"event_twitter"`
	Description           string         `yaml:"description"`
	GaTrackingID          string         `yaml:"ga_tracking_id"`
	Startdate             string         `yaml:"startdate"`
	Enddate               string         `yaml:"enddate"`
	CfpDateStart          string         `yaml:"cfp_date_start"`
	CfpDateEnd            string         `yaml:"cfp_date_end"`
	CfpDateAnnounce       string         `yaml:"cfp_date_announce"`
	CfpOpen               string         `yaml:"cfp_open"`
	CfpLink               string         `yaml:"cfp_link"`
	RegistrationDateStart string         `yaml:"registration_date_start"`
	RegistrationDateEnd   string         `yaml:"registration_date_end"`
	RegistrationClosed    string         `yaml:"registration_closed"`
	RegistrationLink      string         `yaml:"registration_link"`
	Coordinates           string         `yaml:"coordinates"`
	Location              string         `yaml:"location"`
	LocationAddress       string         `yaml:"location_address"`
	NavElements           []NavElement   `yaml:"nav_elements"`
	TeamMembers           []Organizer    `yaml:"team_members"`
	OrganizerEmail        string         `yaml:"organizer_email"`
	ProposalEmail         string         `yaml:"proposal_email"`
	Sponsors              []EventSponsor `yaml:"sponsors"`
	SponsorsAccepted      string         `yaml:"sponsors_accepted"`
	SponsorLevels         []SponsorLevel `yaml:"sponsor_levels"`
}

//TODO: Does NavElement need to be exported?

// NavElement represents a navigation element for the event
type NavElement struct {
	Name string `yaml:"name"`
}

//TODO: Does EventSponsor need to be exported?

// EventSponsor represents a sponsor for an event.
type EventSponsor struct {
	ID    string `yaml:"id"`
	Level string `yaml:"level"`
}

// SponsorLevel represents a level of sponsorship for an event
type SponsorLevel struct {
	ID    string `yaml:"id"`
	Label string `yaml:"label"`
	Max   int    `yaml:"max,omitempty"`
}
