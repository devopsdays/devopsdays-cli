package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createOrganizerCmd represents the "create organizer" command
var createOrganizerCmd = &cobra.Command{
	Use:   "organizer",
	Short: "Creates a new organizer for an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		createOrganizer()
	},
}

// editOrganizerCmd represents the "edit organizer" command
var editOrganizerCmd = &cobra.Command{
	Use:   "organizer",
	Short: "Edit an organizer for an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		editOrganizerFake()
	},
}

// removeOrganizerCmd represents the "Organizer remove" command
var removeOrganizerCmd = &cobra.Command{
	Use:   "organizer",
	Short: "Remove an organizer from an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		removeOrganizer()
	},
}

// showOrganizerCmd represents the "Organizer show" command
var showOrganizerCmd = &cobra.Command{
	Use:   "organizer",
	Short: "Show an organizer from an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		showOrganizer()
	},
}

func init() {
	createCmd.AddCommand(createOrganizerCmd)
	editCmd.AddCommand(editOrganizerCmd)
	removeCmd.AddCommand(removeOrganizerCmd)
	showCmd.AddCommand(showOrganizerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// OrganizerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// OrganizerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// Main functions go down here

func createOrganizer() {
	fmt.Println("You would have created a new Organizer if this happened")
}

func editOrganizerFake() {
	fmt.Println("You would have edited a Organizer if this happened")
}

func removeOrganizer() {
	fmt.Println("You would have removed a Organizer if this happened")
}

func showOrganizer() {
	fmt.Println("You would have shown a Organizer if this happened")
}
