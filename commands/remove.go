package commands

import (
	"github.com/spf13/cobra"
)

// addCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [item]",
	Short: "Remove items from an event, a talk, or a program",
	Long:  `Use this to remove sponsors from events, etc.`,
}

func init() {
	RootCmd.AddCommand(removeCmd)

}
