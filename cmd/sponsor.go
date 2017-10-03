package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/nfnt/resize"
	"github.com/spf13/cobra"
)

// sponsorCmd represents the sponsor command
var sponsorCmd = &cobra.Command{
	Use:   "sponsor [name]",
	Short: "Create a sponsor",
	Long: `Create a new sponsor file, and optionally add the sponsor's image.
The name argument must not contain any spaces.`,
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
		addSponsorFake() //TODO: This is the fake one
	},
}

// createSponsorCmd represents the "create sponsor" command
var createSponsorCmd = &cobra.Command{
	Use:   "sponsor",
	Short: "Create a new sponsor",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		createSponsor()
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

func addSponsorFake() {
	fmt.Println("You would have added a sponsor if this happened")
}

func createSponsor() {
	fmt.Println("You would have created a new sponsor if this happened")
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

func addSponsor(sponsor string) (err error) {

	reader := bufio.NewReader(os.Stdin)
	if sponsor == "" {
		fmt.Println("Enter the sponsor's name. It must not contain any spaces: ")
		sponsor, _ = reader.ReadString('\n')
	}
	// Check if the sponsor exists already
	if checkSponsor(sponsor) == true {
		return errors.New("Sponsor already exists. Try adding it again, perhaps appending '-YYYY'\nFor example, 'chef-2017'")
	}
	// prompt for the path to the sponsor image file
	fmt.Println("Optional: Enter the path to the sponsor's image. It must be the full path. For example: `/Users/mattstratton/chef.png`. Enter return to add the sponsor image manually later.")
	sponsorImage, _ := reader.ReadString('\n')
	if sponsorImage == "\n" {
		fmt.Println("No sponsor image entered. Be sure to copy it to the path ", sponsorImagePath(webdir, sponsor), "later.")
	} else {

		if sponsorImage = strings.TrimSpace(sponsorImage); checkSponsorImage(sponsorImage) == false {
			return errors.New("Sponsor image not found.")
		}
	}

	// prompt for sponsor's name
	fmt.Println("Enter the sponsor's full name. For example: `Chef Software, Inc`")
	sponsorName, _ := reader.ReadString('\n')
	if sponsorName == "\n" {
		return errors.New("Sponsor Name is required.")
	}
	fmt.Println(sponsorName)

	// prompt for sponsor URL
	fmt.Println("Enter the sponsor's URL. It must include 'http://' or 'https://'. For example: `https://www.chef.io`")
	sponsorUrl, _ := reader.ReadString('\n')
	if sponsorUrl == "\n" {
		return errors.New("Sponsor URL is required.")
	}

	// write sponsor YAML file and copy image from path to proper destination
	createSponsorFile(sponsor, sponsorName, sponsorUrl)
	fmt.Println("Sponsor created for ", sponsorName)
	if sponsorImage != "\n" {
		resizeSponsorImage(strings.TrimSpace(sponsorImage), sponsorImagePath(webdir, sponsor))
	} else {
		fmt.Println("Don't forget to place the sponsor image at ", sponsorImagePath(webdir, sponsor))
	}
	return
}

func createSponsorFile(sponsor, sponsorName, sponsorUrl string) (string, error) {
	// find a rice.Box
	templateBox, err := rice.FindBox("../templates")
	if err != nil {
		log.Fatal(err)
	}
	// get file contents as string
	templateString, err := templateBox.String("sponsor.yml.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	// t := template.Must(template.New("sponsor.yml.tmpl").ParseFiles("templates/sponsor.yml.tmpl"))
	t, err := template.New("sponsor.yml").Parse(templateString)
	data := struct {
		Name string
		Url  string
	}{
		strings.TrimSpace(sponsorName),
		strings.TrimSpace(sponsorUrl),
	}
	f, err := os.Create(sponsorDataPath(webdir, sponsor))
	if err != nil {
		return "", err
	}
	defer f.Close()
	t.Execute(f, data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created sponsor file for", sponsor, "at", sponsorDataPath(webdir, sponsor))
	}
	return sponsor, nil

}

// checkSponsor takes in one argument, the name of a sponsor, and returns true if the sponsor already exists.
func checkSponsor(sponsor string) bool {
	fmt.Println(sponsorDataPath(webdir, sponsor))
	if _, err := os.Stat(sponsorDataPath(webdir, sponsor)); err == nil {
		return true
	}
	return false

}

func checkSponsorImage(path string) bool {
	fmt.Println(path)
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false

}

func sponsorDataPath(webdir, sponsor string) (sponsorDataPath string) {
	s := []string{strings.TrimSpace(sponsor), ".yml"}
	// sponsorDataPath = strings.Join(s, "")
	sponsorDataPath = filepath.Join(webdir, "data", "sponsors", strings.Join(s, ""))
	return sponsorDataPath
}

func sponsorImagePath(webdir, sponsor string) (sponsorImagePath string) {
	s := []string{webdir, "/static/img/sponsors/", strings.TrimSpace(sponsor), ".png"}
	sponsorImagePath = strings.Join(s, "")
	return sponsorImagePath
}

func resizeSponsorImage(srcPath, destPath string) {
	fmt.Println("Resizing image")
	file, err := os.Open(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(600, 0, img, resize.Lanczos3)

	out, err := os.Create(destPath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	png.Encode(out, m)
}
