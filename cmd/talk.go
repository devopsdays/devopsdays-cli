package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// talkCmd represents the talk command
var talkCmd = &cobra.Command{
	Use:   "talk",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// addTalkCmd represents the "talk add" command
var addTalkCmd = &cobra.Command{
	Use:   "talk",
	Short: "Add a talk to an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("addTalk called")
	},
}

// editTalkCmd represents the "talk edi" command
var editTalkCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("addTalk called")
	},
}

// removeTalkCmd represents the "talk remove" command
var removeTalkCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("removeTalk called")
	},
}

func init() {
	RootCmd.AddCommand(talkCmd)
	addCmd.AddCommand(addTalkCmd)
	talkCmd.AddCommand(editTalkCmd)
	talkCmd.AddCommand(removeTalkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// talkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// talkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
