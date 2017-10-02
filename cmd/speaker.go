package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addSpeakerCmd represents the "speaker add" command
var addSpeakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Add a speaker to an existing talk",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		addSpeaker()
	},
}

// createSpeakerCmd represents the "speaker create" command
var createSpeakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Creates a new speaker for an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		createSpeaker()
	},
}

// editSpeakerCmd represents the "speaker edi" command
var editSpeakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Edit an event's speaker",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		editSpeaker()
	},
}

// removeSpeakerCmd represents the "speaker remove" command
var removeSpeakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Remove a speaker from an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		removeSpeaker()
	},
}

// showSpeakerCmd represents the "speaker show" command
var showSpeakerCmd = &cobra.Command{
	Use:   "speaker",
	Short: "Show a speaker from an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

func createSpeaker() {
	fmt.Println("You would have created a new speaker if this happened")
}

func editSpeaker() {
	fmt.Println("You would have edited a speaker if this happened")
}

func removeSpeaker() {
	fmt.Println("You would have removed a speaker if this happened")
}

func showSpeaker() {
	fmt.Println("You would have shown a speaker if this happened")
}
