// Package create provides functions to create new content.
package create

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"text/template"

	"github.com/devopsdays/devopsdays-cli/helpers"
	"github.com/devopsdays/devopsdays-cli/model"
	"github.com/fatih/color"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

const tmpl = `+++
Title = "{{ .Title }}"
type = "speaker"
{{ with .Website }}website = "{{ . }}"{{ end }}
{{ with .Twitter }}twitter = "{{ . }}"{{ end }}
{{ with .Facebook }}facebook = "{{ . }}"{{ end }}
{{ with .Linkedin }}linkedin = "{{ . }}"{{ end }}
{{ with .Github }}github = "{{ . }}"{{ end }}
{{ with .Gitlab }}gitlab = "{{ . }}"{{ end }}
{{ with .ImagePath }}image = "{{ . }}"{{ end }}
+++
{{ with .Bio }}{{.}}{{ end }}
`

// the questions to ask
var qs = []*survey.Question{
	{
		Name: "name",
		Prompt: &survey.Input{
			Message: "What is the speaker's name?",
			Help:    "This is the speaker's full name",
		},
		Validate: survey.Required,
	},
	{
		Name:   "bio",
		Prompt: &survey.Editor{Message: "Enter the speaker's bio [Enter to launch editor]"},
	},
	{
		Name: "twitter",
		Prompt: &survey.Input{
			Message: "What is the speaker's Twitter? [optional]",
			Help:    "Twitter username should not include the @ symbol",
		},
		Validate: func(val interface{}) error {
			if str, _ := val.(string); (str != "") && (helpers.ValidateField(str, "twitter") == false) {
				return errors.New("Please enter a valid Twitter handle. It should not have the @ symbol.")
			}
			return nil
		},
	},
	{
		Name: "facebook",
		Prompt: &survey.Input{
			Message: "What is the speaker's Facebook URL? [optional]",
			Help:    "This should be the full URL to the speaker's Facebook page. Example: https://www.facebook.com/matt.stratton",
		},
		Validate: func(val interface{}) error {
			if str, _ := val.(string); (str != "") && (helpers.ValidateField(str, "facebook") == false) {
				return errors.New("Please enter a valid Facebook URL. It should be a URL.")
			}
			return nil
		},
	},
	{
		Name: "linkedin",
		Prompt: &survey.Input{
			Message: "What is the speaker's LinkedIn URL? [optional]",
			Help:    "This should be the full URL to the speaker's LinkedIn profile. Example: https://www.linkedin.com/in/mattstratton/",
		},
		Validate: func(val interface{}) error {
			if str, _ := val.(string); (str != "") && (helpers.ValidateField(str, "linkedin")) == false {
				return errors.New("Please enter a valid LinkedIn URL. It should be a URL.")
			}
			return nil
		},
	},
	{
		Name: "github",
		Prompt: &survey.Input{
			Message: "What is the speaker's GitHub username?",
			Help:    "This should be the username, not the URL. Example: mattstratton",
		},
		Validate: func(val interface{}) error {
			if str, _ := val.(string); (str != "") && (helpers.ValidateField(str, "github")) == false {
				return errors.New("Please enter a valid GitHub username. It should be the username, not URL.")
			}
			return nil
		},
	},
	{
		Name: "gitlab",
		Prompt: &survey.Input{
			Message: "What is the speaker's GitLab username? [optional]",
			Help:    "This should be the username, not the URL. Example: mattstratton",
		},
		Validate: func(val interface{}) error {
			if str, _ := val.(string); (str != "") && (helpers.ValidateField(str, "gitlab")) == false {
				return errors.New("Please enter a valid GitLab username. It should be the username, not URL.")
			}
			return nil
		},
	},
	{
		Name: "imagepath",
		Prompt: &survey.Input{
			Message: "Enter the path to the speaker's image. [optional]",
			Help:    "Path to speaker image. Must be a PNG or JPG file. Example: /Users/mattstratton/Pictures/matt-stratton.png",
		},
		Validate: func(val interface{}) error {
			str, _ := val.(string)
			if str != "" {
				if _, err := os.Stat(str); err != nil {
					return errors.New("File not found.")
				}
			}

			return nil
		},
	},
}

// Speaker takes input from the user to create a new speaker
func Speaker(speakerName, city, year string) (err error) {

	// var imagePath string

	answers := struct {
		Name      string
		Bio       string
		Website   string
		Twitter   string
		Facebook  string
		Linkedin  string
		Github    string
		Gitlab    string
		ImagePath string
	}{}

	if city == "" {
		prompt := &survey.Input{
			Message: "Enter the city name:",
		}
		survey.AskOne(prompt, &city, survey.Required)
	}

	if year == "" {
		prompt := &survey.Input{
			Message: "Enter the year:",
		}
		survey.AskOne(prompt, &year, survey.Required)
	}

	surveyErr := survey.Ask(qs, &answers)
	if surveyErr != nil {
		fmt.Println(surveyErr.Error())
		return
	}

	if answers.ImagePath != "" {
		answers.ImagePath = SpeakerImage(answers.ImagePath, helpers.NameClean(answers.Name), city, year)
	}

	mySpeaker := model.Speaker{
		Name:      helpers.NameClean(answers.Name),
		Title:     answers.Name,
		Website:   answers.Website,
		Twitter:   answers.Twitter,
		Facebook:  answers.Facebook,
		Linkedin:  answers.Linkedin,
		Github:    answers.Github,
		Gitlab:    answers.Gitlab,
		ImagePath: answers.ImagePath,
		Bio:       answers.Bio,
	}

	NewSpeaker(mySpeaker, city, year)

	return
}

// NewSpeaker takes in a constructed Speaker type and generates the stuff
func NewSpeaker(speaker model.Speaker, city string, year string) (err error) {

	cleanName := helpers.NameClean(speaker.Name)
	t := template.New("Speaker template")

	t, err = t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	// err = t.Execute(os.Stdout, speaker)
	// if err != nil {
	// 	log.Fatal("Execute: ", err)
	// 	return
	// }

	s := []string{strings.TrimSpace(cleanName), ".md"}
	f, err := os.Create(filepath.Join(helpers.EventContentPath(city, year), "speakers", strings.Join(s, "")))
	if err != nil {
		return err
	}
	defer f.Close()
	t.Execute(f, speaker)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(color.Output, "\n\n\nCreated speaker file for %s\n", color.GreenString(speaker.Title))
		fmt.Fprintf(color.Output, "at %s\n\n\n", color.BlueString(filepath.Join(helpers.EventContentPath(city, year), "speakers", strings.Join(s, ""))))
	}
	return
}

// SpeakerImage takes in a path to an image and resizes it to the proper dimensions and copies it to the destination
func SpeakerImage(srcPath, speaker, city, year string) (imageFile string) {

	var eventStaticPath string
	var err error
	eventStaticPath, err = helpers.EventStaticPath(city, year)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(filepath.Join(eventStaticPath, "speakers"), 0777); err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`\.[^.]+$`)
	ext := strings.ToLower(re.FindString(srcPath))
	switch ext {
	case ".jpg":
		s := []string{strings.TrimSpace(speaker), ".jpg"}
		destPath := filepath.Join(eventStaticPath, "speakers", strings.Join(s, ""))
		helpers.ResizeImage(srcPath, destPath, "jpg", 600, 600)
		return strings.Join(s, "")
	case ".jpeg":
		s := []string{strings.TrimSpace(speaker), ".jpg"}
		destPath := filepath.Join(eventStaticPath, "speakers", strings.Join(s, ""))
		helpers.ResizeImage(srcPath, destPath, "jpg", 600, 600)
		return strings.Join(s, "")
	case ".png":
		s := []string{strings.TrimSpace(speaker), ".png"}
		destPath := filepath.Join(eventStaticPath, "speakers", strings.Join(s, ""))
		helpers.ResizeImage(srcPath, destPath, "png", 600, 600)
		return strings.Join(s, "")
	}
	return "busted"
}
