package cmd

import (
	"fmt"

	"github.com/devopsdays/devopsdays-cli/event"
	"github.com/spf13/cobra"
)

// eventCmd represents the "event" command
var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// TODO: Work your own magic here
	// 	fmt.Println("speaker called")
	// },
}

// createEventCmd represents the "create event" command
var createEventCmd = &cobra.Command{
	Use:   "event create",
	Short: "Create a new event",
	Long: `Create a new event.
`,
	Example: `  devopsdays-cli create event
  devopsdays-cli create event -c New York --year 2017`,

	Run: func(cmd *cobra.Command, args []string) {
		event.CreateEvent(City, Year)
	},
}

// editEventCmd represents the "edit event" command
var editEventCmd = &cobra.Command{
	Use:   "event",
	Short: "Edit an existing event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		showEvent()

	},
}

// showEventCmd represents the "show event" command
var showEventCmd = &cobra.Command{
	Use:   "event",
	Short: "Show a event from an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		showEvent()
	},
}

func init() {
	// RootCmd.AddCommand(eventCmd)
	createCmd.AddCommand(createEventCmd)
	editCmd.AddCommand(editEventCmd)
	showCmd.AddCommand(showEventCmd)

	createEventCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	createEventCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")
	editEventCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	editEventCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")
	showEventCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	showEventCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")

}

// Main functions go here

func showEvent() {
	fmt.Println("You would have shown an event if this happened")
}

func editEvent() {
	fmt.Println("You would have edited an event if this happened")

}
