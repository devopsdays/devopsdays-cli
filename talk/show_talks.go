package talk

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/color"
	"gopkg.in/AlecAivazis/survey.v1"
)

func ShowTalks(city, year string) (err error) {
	var selection string

	myTalks, err := loadTalks(city, year)

	var options2 []string

	for _, d := range myTalks {
		options2 = append(options2, d.Title)
	}
	options2 = append(options2, "Return to Main Menu")
	for selection != "Return to Main Menu" {
		prompt := &survey.Select{
			Message: "Select a talk:",
			Options: options2,
		}
		survey.AskOne(prompt, &selection, nil)
		if selection == "Return to Main Menu" {
			return
		}
		var myTalk *Talk
		for _, d := range myTalks {
			if d.Title == selection {
				myTalk = d
				break
			}
		}
		s := reflect.ValueOf(myTalk).Elem()
		typeOfT := s.Type()
		fmt.Println()
		color.Cyan("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			if (typeOfT.Field(i).Name != "Name") && (f.Interface() != "") {
				if f.Type() != reflect.TypeOf("") {
					continue
				}
				str := f.Interface().(string)
				str = strings.TrimSpace(str)
				fmt.Fprintf(color.Output, "%s: %s\n", color.CyanString(typeOfT.Field(i).Name), color.GreenString(str))
			}
		}
		color.Cyan("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
		fmt.Println()

	}
	return
}
