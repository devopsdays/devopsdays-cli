package commands

import (
	"fmt"

	speaker "github.com/devopsdays/devopsdays-cli/speaker"
	"github.com/spf13/cobra"
)

// addSpeakerCmd represents the "speaker add" command
var addSpeakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Add a speaker to an existing talk",
	Long: `Add a speaker to an existing talk.
You can provide the speaker's name as an argument to this command, but it must be the "cleaned" name.
		`,
	Example: `  devopsdays-cli add speaker george-bluth
  devopsdays-cli add speaker --city new-york
  devopsdays-cli add speaker george-bluth -c "New York" --year "2017"`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addSpeaker()
	},
}

// createSpeakerCmd represents the "speaker create" command
var createSpeakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Creates a new speaker for an event",
	Long: `Creates a new speaker for an event.
	`,
	Example: `  devopsdays-cli create speaker
  devopsdays-cli create speaker --city new-york
  devopsdays-cli create speaker -c "New York" --year "2017"`,

	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createSpeakerPrompt(City, Year)
	},
}

// editSpeakerCmd represents the "speaker edi" command
var editSpeakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Edit an event's speaker",
	Long: `Edit an event's speaker.
You can provide the speaker's name as an argument to this command, but it must be the "cleaned" name.
`,
	Example: `  devopsdays-cli edit speaker george-bluth
  devopsdays-cli edit speaker --city new-york
  devopsdays-cli edit speaker george-bluth -c "New York" --year "2017"`,
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
	Long: `Remove a speaker from an event.
You can provide the speaker's name as an argument to this command, but it must be the "cleaned" name.
	`,
	Example: `  devopsdays-cli remove speaker george-bluth
  devopsdays-cli remove speaker --city new-york
  devopsdays-cli remove speaker george-bluth -c "New York" --year "2017"`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		removeSpeaker()
	},
}

// showSpeakerCmd represents the "speaker show" command
var showSpeakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Show a speaker from an event",
	Long: `Show a speaker from an event.
		`,
	Example: `  devopsdays-cli show speaker george-bluth
  devopsdays-cli show speaker --city new-york --year 2017 --all
  devopsdays-cli show speaker george-bluth -c "New York" --year "2017"`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		showSpeakerPrompt(City, Year)
	},
}

func init() {
	addCmd.AddCommand(addSpeakerCmd)
	createCmd.AddCommand(createSpeakerCmd)
	editCmd.AddCommand(editSpeakerCmd)
	removeCmd.AddCommand(removeSpeakerCmd)
	showCmd.AddCommand(showSpeakerCmd)

	addSpeakerCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	addSpeakerCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")
	createSpeakerCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	createSpeakerCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")
	editSpeakerCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	editSpeakerCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")
	removeSpeakerCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	removeSpeakerCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")
	showSpeakerCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	showSpeakerCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")
	showSpeakerCmd.Flags().BoolVarP(&All, "all", "a", false, "show all")

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
	fmt.Println("Showing a speaker")
	mySpeaker, _ := speaker.GetSpeakerInfo("fluttershy.md", "ponyville", "2017")
	fmt.Println(mySpeaker)
}
