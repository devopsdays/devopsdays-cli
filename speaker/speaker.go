// Package speaker provides functions to add, create, edit, delete, and show speakers
package speaker

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

//  Prompts for a new speaker
var qsCreateSpeaker = []*survey.Question{
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
			Message: "What is the speaker's GitHub username? [optional]",
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

// CreateSpeaker takes input from the user to create a new speaker
func CreateSpeaker(speakerName, city, year string) (err error) {

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
		Talk      string
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

	surveyErr := survey.Ask(qsCreateSpeaker, &answers)
	if surveyErr != nil {
		fmt.Println(surveyErr.Error())
		return
	}

	name := false
	prompt := &survey.Confirm{
		Message: "Do you want to add this speaker to an existing talk?",
	}
	survey.AskOne(prompt, &name, nil)

	talk := ""
	if name == true {
		prompt := &survey.Select{
			Message: "Choose a talk:",
			Options: helpers.GetTalks(city, year),
		}
		survey.AskOne(prompt, &talk, nil)
		color.Yellow("NOT IMPLEMENTED")
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

	t, err = t.Parse(speakerTmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	if err := os.MkdirAll(filepath.Join(helpers.EventContentPath(city, year), "speakers"), 0777); err != nil {
		log.Fatal(err)
	}
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
