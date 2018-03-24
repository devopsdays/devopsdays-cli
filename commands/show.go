package commands

import (
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show [item]",
	Short: "Show details about various items",
	Long:  `The show command displays information about events, sponsors, talks, configuration, etc.`,
}

func init() {
	RootCmd.AddCommand(showCmd)

}
