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
	"github.com/devopsdays/devopsdays-cli/helpers/paths"
	"github.com/devopsdays/devopsdays-cli/images"
	"github.com/devopsdays/devopsdays-cli/model"
	"github.com/devopsdays/devopsdays-cli/names"
	"github.com/devopsdays/devopsdays-cli/talk"
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
		prompt := &survey.Select{
			Message: "Choose a talk:",
			Options: talk.GetTalks(city, year),
		}
		survey.AskOne(prompt, &t, nil)
		color.Yellow("NOT IMPLEMENTED")
	}

	if answers.ImagePath != "" {
		answers.ImagePath = SpeakerImage(answers.ImagePath, names.NameClean(answers.Name), city, year)
	}

	mySpeaker := model.Speaker{
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

// NewSpeaker takes in a constructed Speaker type and generates the stuff
func NewSpeaker(speaker model.Speaker, city string, year string) (err error) {

	cleanName := names.NameClean(speaker.Name)
	t := template.New("Speaker template")

	t, err = t.Parse(speakerTmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	if err := os.MkdirAll(filepath.Join(paths.EventContentPath(city, year), "speakers"), 0777); err != nil {
		log.Fatal(err)
	}
	s := []string{strings.TrimSpace(cleanName), ".md"}
	f, err := os.Create(filepath.Join(paths.EventContentPath(city, year), "speakers", strings.Join(s, "")))
	if err != nil {
		return err
	}
	defer f.Close()
	t.Execute(f, speaker)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(color.Output, "\n\n\nCreated speaker file for %s\n", color.GreenString(speaker.Title))
		fmt.Fprintf(color.Output, "at %s\n\n\n", color.BlueString(filepath.Join(paths.EventContentPath(city, year), "speakers", strings.Join(s, ""))))
	}
	return
}

// SpeakerImage takes in a path to an image and resizes it to the proper dimensions and copies it to the destination
func SpeakerImage(srcPath, speaker, city, year string) (imageFile string) {

	var eventStaticPath string
	var err error
	eventStaticPath, err = paths.EventStaticPath(city, year)
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
		images.ResizeImage(srcPath, destPath, "jpg", 600, 600)
		return strings.Join(s, "")
	case ".jpeg":
		s := []string{strings.TrimSpace(speaker), ".jpg"}
		destPath := filepath.Join(eventStaticPath, "speakers", strings.Join(s, ""))
		images.ResizeImage(srcPath, destPath, "jpg", 600, 600)
		return strings.Join(s, "")
	case ".png":
		s := []string{strings.TrimSpace(speaker), ".png"}
		destPath := filepath.Join(eventStaticPath, "speakers", strings.Join(s, ""))
		images.ResizeImage(srcPath, destPath, "png", 600, 600)
		return strings.Join(s, "")
	}
	return "busted"
}

func ShowSpeakers(city, year string) (err error) {
	var selection string

	speakerList, _ := GetSpeakers(city, year)
	options2, _ := listSpeakerNames(speakerList, city, year)

	options2 = append(options2, "Return to Main Menu")
	for selection != "Return to Main Menu" {
		prompt := &survey.Select{
			Message: "Select a speaker:",
			Options: options2,
		}
		survey.AskOne(prompt, &selection, nil)
		if selection == "Return to Main Menu" {
			return
		}
		speakerFileName := strings.Join([]string{strings.TrimSpace(names.NameClean(selection)), ".md"}, "")

		var mySpeaker model.Speaker
		mySpeaker, err = GetSpeakerInfo(speakerFileName, city, year)
		fmt.Println()
		color.Cyan("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
		fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("Name: "), color.GreenString(mySpeaker.Title))

		if mySpeaker.Website != "" {
			fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("WebsiteName: "), color.GreenString(mySpeaker.Website))
		}
		if mySpeaker.Twitter != "" {
			fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("Twitter: "), color.GreenString(fmt.Sprintf("@%s", mySpeaker.Twitter)))
		}
		if mySpeaker.Facebook != "" {
			fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("Facebook: "), color.GreenString(mySpeaker.Facebook))
		}
		if mySpeaker.Linkedin != "" {
			fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("LinkedIn: "), color.GreenString(mySpeaker.Linkedin))
		}
		if mySpeaker.Github != "" {
			fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("GitHub: "), color.GreenString(mySpeaker.Github))
		}
		if mySpeaker.Gitlab != "" {
			fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("GitLab: "), color.GreenString(mySpeaker.Gitlab))
		}
		if mySpeaker.Bio != "" {
			fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("Bio: "), color.GreenString(mySpeaker.Bio))
		}
		color.Cyan("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
		fmt.Println()
	}
	return
}

func listSpeakerNames(speakers []string, city string, year string) (speakerFullNames []string, err error) {
	for _, f := range speakers {
		var mySpeaker model.Speaker
		mySpeaker, err = GetSpeakerInfo(f, city, year)
		speakerFullNames = append(speakerFullNames, mySpeaker.Title)
	}
	return
}
