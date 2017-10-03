package cmd

import (
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [item]",
	Short: "Edit an existing item",
	Long:  `Use this to edit events, talks, organizers, etc.`,
}

func init() {
	RootCmd.AddCommand(editCmd)

}
