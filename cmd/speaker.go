package cmd

import (
	"fmt"
	"log"

	"github.com/devopsdays/devopsdays-cli/create"
	"github.com/devopsdays/devopsdays-cli/helpers"
	"github.com/devopsdays/devopsdays-cli/model"
	"github.com/spf13/cobra"
	"github.com/tcnksm/go-input"
)

// addSpeakerCmd represents the "speaker add" command
var addSpeakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Add a speaker to an existing talk",
	Long: `Add a speaker to an existing talk

You can provide the speaker's name as an argument to this command, but it must be in quotes. For example:
	devopsdays-cli create speaker "George Bluth"`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		addSpeaker()
	},
}

// createSpeakerCmd represents the "speaker create" command
var createSpeakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Creates a new speaker for an event",
	Long: `Creates a new speaker for an event

You can provide the speaker's name as an argument to this command, but it must be in quotes. For example:
devopsdays-cli create speaker "George Bluth"`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			createSpeaker(args[0], "", "")
		} else {
			createSpeaker("", "", "")
		}
	},
}

// editSpeakerCmd represents the "speaker edi" command
var editSpeakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Edit an event's speaker",
	Long: `Edit an event's speaker

You can provide the speaker's name as an argument to this command, but it must be the "cleaned" name. For example:

	devopsdays-cli edit speaker george-bluth`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			editSpeaker(args[0], "", "")
		} else {
			editSpeaker("", "", "")
		}
	},
}

// removeSpeakerCmd represents the "speaker remove" command
var removeSpeakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Remove a speaker from an event",
	Long: `Remove a speaker from an event
You can provide the speaker's name as an argument to this command, but it must be the "cleaned" name. For example:

		devopsdays-cli remove speaker george-bluth`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		removeSpeaker()
	},
}

// showSpeakerCmd represents the "speaker show" command
var showSpeakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Show a speaker from an event",
	Long: `Show a speaker from an event
You can provide the speaker's name as an argument to this command, but it must be the "cleaned" name. For example:

		devopsdays-cli show speaker george-bluth`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		showSpeaker()
	},
}

func init() {
	addCmd.AddCommand(addSpeakerCmd)
	createCmd.AddCommand(createSpeakerCmd)
	editCmd.AddCommand(editSpeakerCmd)
	removeCmd.AddCommand(removeSpeakerCmd)
	showCmd.AddCommand(showSpeakerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// speakerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// speakerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// Main functions go down here

func addSpeaker() {
	fmt.Println("You would have added a speaker to a talk if this happened")
}

func createSpeaker(speakerName, city, year string) (err error) {
	// debug
	fmt.Println("Creating new speaker")
	fmt.Println("Value of city is")
	fmt.Println(cityFlag)
	// actual stuff
	ui := &input.UI{}

	if city == "" {
		cityName, err := ui.Ask("City", &input.Options{
			// Read the default val from env var
			Required:  true,
			Loop:      true,
			HideOrder: true,
		})
		if err != nil {
			log.Fatal(err)
		}
		city = cityName
	}

	if year == "" {
		yearName, err := ui.Ask("Year", &input.Options{
			// Read the default val from env var
			Required:  true,
			Loop:      true,
			HideOrder: true,
		})
		if err != nil {
			log.Fatal(err)
		}
		year = yearName
	}

	name, err := ui.Ask("Speaker Name", &input.Options{
		Required:  true,
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	website, err := ui.Ask("Website (optional)", &input.Options{
		HideOrder: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	twitter, err := ui.Ask("Twitter (optional)", &input.Options{
		HideOrder: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	facebook, err := ui.Ask("Facebook (optional)", &input.Options{
		HideOrder: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	linkedin, err := ui.Ask("LinkedIn (optional)", &input.Options{
		HideOrder: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	github, err := ui.Ask("GitHub (optional)", &input.Options{
		HideOrder: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	gitlab, err := ui.Ask("GitLab (optional)", &input.Options{
		HideOrder: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	mySpeaker := model.Speaker{
		Name:      helpers.NameClean(name),
		Title:     name,
		Website:   website,
		Twitter:   twitter,
		Facebook:  facebook,
		Linkedin:  linkedin,
		Github:    github,
		Gitlab:    gitlab,
		ImagePath: "",
	}

	create.NewSpeaker(mySpeaker, "ponyville", "2017")

	return
}

func editSpeaker(speakerName, city, year string) {
	fmt.Println("You would have edited a speaker if this happened")
}

func removeSpeaker() {
	fmt.Println("You would have removed a speaker if this happened")
}

func showSpeaker() {
	fmt.Println("You would have shown a speaker if this happened")
}
