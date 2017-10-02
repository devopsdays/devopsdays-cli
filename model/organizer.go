package model

// Organizer defines a devopsdays event organizer
type Organizer struct {
	Name     string `yaml:"name"`
	Twitter  string `yaml:"twitter,omitempty"`
	Employer string `yaml:"employer,omitempty"`
	Github   string `yaml:"github,omitempty"`
	Facebook string `yaml:"facebook,omitempty"`
	Linkedin string `yaml:"linkedin,omitempty"`
	Website  string `yaml:"website,omitempty"`
	Image    string `yaml:"image,omitempty"`
	Bio      string `yaml:"bio,omitempty"`
}
