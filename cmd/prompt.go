package cmd

import (
	"github.com/AlecAivazis/survey"
	"github.com/devopsdays/devopsdays-cli/speaker"
)

func createSpeakerPrompt(city, year string) (err error) {
	var exitCode = true

	for exitCode {
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
		speaker.CreateSpeaker("", city, year)
		prompt := &survey.Confirm{
			Message: "Do you want to add another speaker?",
		}
		survey.AskOne(prompt, &exitCode, nil)
	}
	return
}
