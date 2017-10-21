// Package talk includes the functionality to add, create, edit, remove, and show talks.
// It also includes supporting and helper functions that are talk-releated.
package talk

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/AlecAivazis/survey"
	"github.com/devopsdays/devopsdays-cli/helpers/paths"
	"github.com/devopsdays/devopsdays-cli/model"
	"github.com/fatih/color"
	"github.com/pkg/errors"
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
		var myTalk *model.Talk
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

func loadTalks(city, year string) (talks []*model.Talk, err error) {
	talksDir := filepath.Join(paths.EventContentPath(city, year), "program")
	files, err := ioutil.ReadDir(talksDir)
	if err != nil {
		return nil, errors.Wrap(err, "read directory failed")
	}

	for _, f := range files {
		thisTalk, _ := GetTalkInfo(f.Name(), city, year)
		talks = append(talks, thisTalk)
	}

	return
}
