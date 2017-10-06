package model

// Sponsor represents a sponsor inside a sponsor data file
type Sponsor struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
