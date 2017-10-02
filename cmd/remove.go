package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [thing]",
	Short: "Remove a thing to another thing",
	Long:  `Use this to remove sponsors to events, etc.`,
	// Args:  cobra.MinimumNArgs(1),
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// TODO: Work your own magic here
	// 	fmt.Println("remove called")
	// },
}

func init() {
	RootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
