package cmd

import (
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [thing]",
	Short: "Edit an existing thing",
	Long:  `Use this to edit events, talks, organizers, etc.`,
	// Args:  cobra.MinimumNArgs(1),
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// TODO: Work your own magic here
	// 	fmt.Println("edit called")
	// },
}

func init() {
	RootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
