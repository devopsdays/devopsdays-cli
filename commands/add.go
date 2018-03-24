package commands

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [item]",
	Short: "Add items to talks, programs, or events",
	Long:  `Use this to add sponsors to events, talks to programs, etc.`,
}

func init() {
	RootCmd.AddCommand(addCmd)

}
