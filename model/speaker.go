package model

// Speaker defines a devopsdays event's speaker
type Speaker struct {
	Name      string
	Title     string
	Website   string `toml:"website,omitempty"`
	Twitter   string `toml:"twitter,omitempty"`
	Facebook  string `toml:"facebook,omitempty"`
	Linkedin  string `toml:"linkedin,omitempty"`
	Github    string `toml:"github,omitempty"`
	Gitlab    string `toml:"gitlab,omitempty"`
	ImagePath string `toml:"image,omitempty"`
	Bio       string
}
