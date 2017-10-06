package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sponsorCmd represents the sponsor command
var sponsorCmd = &cobra.Command{
	Use:   "sponsor [name]",
	Short: "Create a sponsor",
	Long: `Create a new sponsor file add the sponsor's image.
The name argument must not contain any spaces.
     `,
	Example: `  devopsdays-cli create sponsor
devopsdays-cli create sponsor bluth-company`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			addSponsor(args[0])
		} else {
			addSponsor("")
		}
	},
}

// addSponsorCmd represents the "add sponsor" command
var addSponsorCmd = &cobra.Command{
	Use:   "sponsor",
	Short: "Add a sponsor to an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		addSponsor("") //TODO: This is the fake one
	},
}

// createSponsorCmd represents the "create sponsor" command
var createSponsorCmd = &cobra.Command{
	Use:   "sponsor [name]",
	Short: "Create a sponsor",
	Long: `Create a new sponsor file add the sponsor's image.
The name argument must not contain any spaces.
     `,
	Example: `  devopsdays-cli create sponsor
  devopsdays-cli create sponsor bluth-company`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			addSponsor(args[0])
		} else {
			addSponsor("")
		}
	},
}

// editSponsorCmd represents the "edit sponsor" command
var editSponsorCmd = &cobra.Command{
	Use:   "sponsor",
	Short: "Create a new sponsor",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly edit a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		editSponsor()
	},
}

// removeSponsorCmd represents the "remove sponsor" command
var removeSponsorCmd = &cobra.Command{
	Use:   "sponsor",
	Short: "Remove a sponsor from an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		removeSponsor() //TODO: This is the fake one
	},
}

// showSponsorCmd represents the "show sponsor" command
var showSponsorCmd = &cobra.Command{
	Use:   "sponsor",
	Short: "Remove a sponsor from an event",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		showSponsor() //TODO: This is the fake one
	},
}

func init() {
	// RootCmd.AddCommand(sponsorCmd)
	addCmd.AddCommand(addSponsorCmd)
	createCmd.AddCommand(createSponsorCmd)
	editCmd.AddCommand(editSponsorCmd)
	removeCmd.AddCommand(removeSponsorCmd)
	showCmd.AddCommand(showSponsorCmd)

	addSponsorCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	addSponsorCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")
	removeSponsorCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	removeSponsorCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")
	showSponsorCmd.Flags().StringVarP(&City, "city", "c", "", "city to use")
	showSponsorCmd.Flags().StringVarP(&Year, "year", "y", "", "year to use")

}

// Main functions go down here

func addSponsor(sponsor string) {
	fmt.Println("You would have added a sponsor if this happened")
}

func editSponsor() {
	fmt.Println("You would have edited an existing sponsor if this happened")
}

func removeSponsor() {
	fmt.Println("You would have removed a sponsor from the event if this happened")
}

func showSponsor() {
	fmt.Println("You would have shown a sponsor if this happened")
}
