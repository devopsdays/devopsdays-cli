package cmd

import (
	"github.com/AlecAivazis/survey"
	"github.com/devopsdays/devopsdays-cli/event"
	"github.com/devopsdays/devopsdays-cli/speaker"
	"github.com/devopsdays/devopsdays-cli/sponsor"
)

func mainPrompt() (err error) {
	var selection string

	for selection != "Quit the application" {
		prompt := &survey.Select{
			Message: "Select an action:",
			Options: []string{
				"Create a new event",
				"Create a new speaker",
				"Create a new sponsor",
				"Show a speaker",
				"Quit the application",
			},
		}
		survey.AskOne(prompt, &selection, nil)
		switch selection {
		case "Create a new event":
			event.CreateEvent("", "")
		case "Create a new speaker":
			createSpeakerPrompt("", "")
		case "Create a new sponsor":
			sponsor.CreateSponsor("")
		case "Show a speaker":
			showSpeakerPrompt("", "")
		}
	}

	return
}

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

func showSpeakerPrompt(city, year string) (err error) {

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
	speaker.ShowSpeakers(city, year)

	return
}
