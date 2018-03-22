package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createProgramCmd represents the "program create" command
var createProgramCmd = &cobra.Command{
	Use:   "program",
	Short: "Create a new program for an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		createProgram()
	},
}

// editProgramCmd represents the "program edi" command
var editProgramCmd = &cobra.Command{
	Use:   "program",
	Short: "Edit an event's program",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		editProgram()
	},
}

// showProgramCmd represents the "program show" command
var showProgramCmd = &cobra.Command{
	Use:   "program",
	Short: "Show an event's program",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		showProgram()
	},
}

func init() {
	createCmd.AddCommand(createProgramCmd)
	editCmd.AddCommand(editProgramCmd)
	showCmd.AddCommand(showProgramCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// programCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// programCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// Main functions go down here

func createProgram() {
	fmt.Println("You would have created a program if this happened")
}

func editProgram() {
	fmt.Println("You would have edited a program if this happened")
}

func showProgram() {
	fmt.Println("You would have shown the program if this happened")
}
