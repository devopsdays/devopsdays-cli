package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/devopsdays/devopsdays-cli/talk"
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
		removeTalk()
	},
}

// showTalkCmd represents the "show talk" command
var showTalkCmd = &cobra.Command{
	Use:   "talk",
	Short: "Show a talk from an event",
	Long: `Show a speaker from an event.
	`,
	Example: `  devopsdays-cli show talk
  devopsdays-cli show talk --city new-york --year 2017 --all
  devopsdays-cli show talk -c "New York" --year "2017"`,
	Run: func(cmd *cobra.Command, args []string) {
		showTalkPrompt(City, Year)
	},
}

func init() {
	addCmd.AddCommand(addTalkCmd)
	createCmd.AddCommand(createTalkCmd)
	editCmd.AddCommand(editTalkCmd)
	removeCmd.AddCommand(removeTalkCmd)
	showCmd.AddCommand(showTalkCmd)

	addTalkCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	addTalkCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")
	createTalkCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	createTalkCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")
	editTalkCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	editTalkCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")
	removeTalkCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	removeTalkCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")
	showTalkCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	showTalkCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")
	showTalkCmd.Flags().BoolVarP(&All, "all", "a", false, "show all NOT IMPLEMENTED")

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
	err := talk.ShowTalks("Ponyville", "2017")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("You would have shown a talk if this happened")
}
