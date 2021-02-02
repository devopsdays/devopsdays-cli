// Package sponsor provides functionality for adding, creating, editing, and showing sponsors.
package sponsor

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	"text/template"

	"github.com/devopsdays/devopsdays-cli/helpers"
	"github.com/devopsdays/devopsdays-cli/images"
	"github.com/devopsdays/devopsdays-cli/model"
	"github.com/fatih/color"

	survey "github.com/AlecAivazis/survey/v2"
)

var qsCreateSponsor = []*survey.Question{

	{
		Name: "fullname",
		Prompt: &survey.Input{
			Message: "What is the sponsors's full name?",
			Help:    "This is the full name of the sponsor. Example: `Chef Software, Inc`",
		},
		Validate: survey.Required,
	},
	{
		Name: "website",
		Prompt: &survey.Input{
			Message: "What is the sponsors's website?",
			Help:    "This should be the full URL to the speaker's website. Example: https://www.arresteddevops.com",
		},
		Validate: func(val interface{}) error {
			if str, _ := val.(string); (str != "") && (helpers.ValidateField(str, "website") == false) {
				return errors.New("Please enter a valid URL.")
			}
			return nil
		},
	},
	{
		Name: "imagepath",
		Prompt: &survey.Input{
			Message: "Enter the path to the sponsor's image.",
			Help:    "Path to sponsor image. Must be a PNG file. Example: /Users/mattstratton/Pictures/arrested-devops.png",
		},
		Validate: func(val interface{}) error {
			str, _ := val.(string)
			if str == "" {
				return errors.New("Sponsor image is required.")
			}

			ret, _ := regexp.MatchString(`[0-9a-z]+\.(png|PNG)`, str)
			if ret != true {
				return errors.New("Sponsor image must be a PNG file")
			}
			if str != "" {
				if _, err := os.Stat(str); err != nil {
					return errors.New("File not found.")
				}
			}

			return nil
		},
	},
}

// CreateSponsor takes input from the user to create a new sponsor
func CreateSponsor(sponsorName string) (err error) {

	answers := struct {
		FullName  string
		Website   string
		ImagePath string
	}{}

	if sponsorName == "" {
		prompt := &survey.Input{
			Message: "What is the sponsors's name, without spaces?",
			Help:    "This is the name of the sponsor. Examples: `chef` or `arrested-devops`",
		}
		survey.AskOne(prompt, &sponsorName, survey.WithValidator(survey.Required))
	}

	if checkSponsor(sponsorName) {
		fmt.Println("This sponsor already exists. If you would like to edit it, please run `devopsdays-cli edit sponsor`")
		return
	}

	surveyErr := survey.Ask(qsCreateSponsor, &answers)
	if surveyErr != nil {
		fmt.Println(surveyErr.Error())
		return
	}

	mySponsor := model.Sponsor{
		Name: answers.FullName,
		URL:  answers.Website,
	}

	newSponsor(mySponsor, sponsorName)
	sponsorLogo(answers.ImagePath, sponsorName)
	return
}

func newSponsor(sponsor model.Sponsor, name string) (err error) {
	t := template.New("Sponsor template")

	t, err = t.Parse(sponsorTmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	f, err := os.Create(sponsorDataPath(name))
	defer f.Close()
	t.Execute(f, sponsor)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(color.Output, "\n\n\nCreated sponsor file for %s\n", color.GreenString(sponsor.Name))
		fmt.Fprintf(color.Output, "at %s\n\n\n", color.BlueString(sponsorDataPath(sponsor.Name)))
	}
	return

}

// sponsorLogo takes in a path the sponsor's  logo, and resizes it and copies it to the proper destination
func sponsorLogo(srcPath, sponsor string) (err error) {
	fmt.Println(sponsorImagePath(sponsor))
	destPath := sponsorImagePath(sponsor)
	images.ResizeImage(srcPath, destPath, "png", 600, 0)

	// @TODO update helpers.ResizeImage to return error code and do something with it here

	return nil
}
