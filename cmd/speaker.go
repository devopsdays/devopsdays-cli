package cmd

import (
	"fmt"

	"github.com/devopsdays/devopsdays-cli/create"
	"github.com/spf13/cobra"
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
			create.CreateSpeaker(args[0], "", "")
		} else {
			create.CreateSpeaker("", "", "")
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

func editSpeaker(speakerName, city, year string) {
	fmt.Println("You would have edited a speaker if this happened")
}

func removeSpeaker() {
	fmt.Println("You would have removed a speaker if this happened")
}

func showSpeaker() {
	fmt.Println("You would have shown a speaker if this happened")
}
