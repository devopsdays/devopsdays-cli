package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addTalkCmd represents the "talk add" command
var addTalkCmd = &cobra.Command{
	Use:   "talk",
	Short: "Add a talk to an event's program",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		createTalk()
	},
}

// createTalkCmd represents the "talk create" command
var createTalkCmd = &cobra.Command{
	Use:   "talk",
	Short: "Create a new talk for an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		createTalk()
	},
}

// editTalkCmd represents the "talk edi" command
var editTalkCmd = &cobra.Command{
	Use:   "talk",
	Short: "Edit an event's talk",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		editTalk()
	},
}

// removeTalkCmd represents the "talk remove" command
var removeTalkCmd = &cobra.Command{
	Use:   "talk",
	Short: "Remove a talk from an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		removeTalk()
	},
}

// showTalkCmd represents the "talk show" command
var showTalkCmd = &cobra.Command{
	Use:   "talk",
	Short: "Show a talk from an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		showTalk()
	},
}

func init() {
	addCmd.AddCommand(createTalkCmd)
	createCmd.AddCommand(createTalkCmd)
	editCmd.AddCommand(editTalkCmd)
	removeCmd.AddCommand(removeTalkCmd)
	showCmd.AddCommand(showTalkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// talkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// talkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// Main functions go down here

func addTalk() {
	fmt.Println("You would have added a talk to a program if this happened")
}

func createTalk() {
	fmt.Println("You would have created a talk if this happened")
}

func editTalk() {
	fmt.Println("You would have edited a talk if this happened")
}

func removeTalk() {
	fmt.Println("You would have removed a talk if this happened")
}

func showTalk() {
	fmt.Println("You would have shown a talk if this happened")
}
