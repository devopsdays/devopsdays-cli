package speaker

import (
	"errors"
	"fmt"
	"os"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/devopsdays/devopsdays-cli/helpers"
	"github.com/devopsdays/devopsdays-cli/names"
	"github.com/devopsdays/devopsdays-cli/talk"
	"github.com/fatih/color"
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
		Name: "website",
		Prompt: &survey.Input{
			Message: "What is the speaker's website? [optional]",
			Help:    "This should be the full URL to the speaker's website. Example: https://www.mattstratton.com",
		},
		Validate: func(val interface{}) error {
			if str, _ := val.(string); (str != "") && (helpers.ValidateField(str, "facebook") == false) {
				return errors.New("Please enter a valid URL.")
			}
			return nil
		},
	},
	{
		Name: "twitter",
		Prompt: &survey.Input{
			Message: "What is the speaker's Twitter? [optional]",
			Help:    "Twitter username can include the @ symbol or not. Examples: '@mattstratton' or 'mattstratton",
		},
		Validate: func(val interface{}) error {
			if str, _ := val.(string); (str != "") && (helpers.ValidateField(str, "twitter") == false) {
				return errors.New("Please enter a valid Twitter handle. It should not have any spaces.")
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

	t := ""
	if name == true {
		myList, _ := talk.GetTalks(city, year)
		prompt := &survey.Select{
			Message: "Choose a talk:",
			Options: myList,
		}
		survey.AskOne(prompt, &t, nil)
		color.Yellow("NOT IMPLEMENTED")
	}

	if answers.ImagePath != "" {
		answers.ImagePath = SpeakerImage(answers.ImagePath, names.NameClean(answers.Name), city, year)
	}

	mySpeaker := Speaker{
		Name:      names.NameClean(answers.Name),
		Title:     answers.Name,
		Website:   answers.Website,
		Twitter:   helpers.TwitterClean(answers.Twitter),
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
