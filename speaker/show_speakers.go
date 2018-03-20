package speaker

import (
	"fmt"
	"strings"

	"github.com/devopsdays/devopsdays-cli/names"
	"github.com/fatih/color"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

func ShowSpeakers(city, year string) (exitCode bool, err error) {
	var selection string

	speakerList, _ := GetSpeakers(city, year)
	options2, _ := listSpeakerNames(speakerList, city, year)

	options2 = append(options2, "Return to Main Menu")
	for selection != "Return to Main Menu" {
		prompt := &survey.Select{
			Message: "Select a speaker:",
			Options: options2,
		}
		survey.AskOne(prompt, &selection, nil)
		if selection == "Return to Main Menu" {
			return true, nil
		}
		speakerFileName := strings.Join([]string{strings.TrimSpace(names.NameClean(selection)), ".md"}, "")

		var mySpeaker Speaker
		mySpeaker, err = GetSpeakerInfo(speakerFileName, city, year)
		fmt.Println()
		color.Cyan("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
		fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("Name: "), color.GreenString(mySpeaker.Title))

		if mySpeaker.Website != "" {
			fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("WebsiteName: "), color.GreenString(mySpeaker.Website))
		}
		if mySpeaker.Twitter != "" {
			fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("Twitter: "), color.GreenString(fmt.Sprintf("@%s", mySpeaker.Twitter)))
		}
		if mySpeaker.Facebook != "" {
			fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("Facebook: "), color.GreenString(mySpeaker.Facebook))
		}
		if mySpeaker.Linkedin != "" {
			fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("LinkedIn: "), color.GreenString(mySpeaker.Linkedin))
		}
		if mySpeaker.Github != "" {
			fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("GitHub: "), color.GreenString(mySpeaker.Github))
		}
		if mySpeaker.Gitlab != "" {
			fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("GitLab: "), color.GreenString(mySpeaker.Gitlab))
		}
		if mySpeaker.Bio != "" {
			fmt.Fprintf(color.Output, "%s %s\n", color.CyanString("Bio: "), color.GreenString(mySpeaker.Bio))
		}
		color.Cyan("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
		fmt.Println()
	}
	return true, nil
}
