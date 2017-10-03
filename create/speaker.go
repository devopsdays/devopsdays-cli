// Package create provides functions to create new content.
package create

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"text/template"

	"github.com/devopsdays/devopsdays-cli/helpers"
	"github.com/devopsdays/devopsdays-cli/model"
	"github.com/tcnksm/go-input"
)

const tmpl = `+++
Title = "{{ .Title }}"
type = "speaker"
{{ with .Website }}website = "{{ . }}"{{ end }}
{{ with .Twitter }}twitter = "{{ . }}"{{ end }}
{{ with .Facebook }}facebook = "{{ . }}"{{ end }}
{{ with .Linkedin }}linkedin = "{{ . }}"{{ end }}
{{ with .Github }}github = "{{ . }}"{{ end }}
{{ with .Gitlab }}gitlab = "{{ . }}"{{ end }}
{{ with .ImagePath }}image = "{{ . }}"{{ end }}
+++
Food-truck SpaceTeam pivot earned media agile big data entrepreneur actionable insight iterate unicorn convergence driven moleskine. User centered design piverate physical computing disrupt moleskine co-working fund pivot. Waterfall is so 2000 and late integrate responsive big data agile piverate affordances. Agile earned media pivot viral engaging thought leader prototype workflow.

`

// CreateSpeaker takes input from the user to create a new speaker
func CreateSpeaker(speakerName, city, year string) (err error) {

	var imagePath string

	ui := &input.UI{}

	if city == "" {
		cityName, err := ui.Ask("City", &input.Options{
			// Read the default val from env var
			Required:  true,
			Loop:      true,
			HideOrder: true,
		})
		if err != nil {
			log.Fatal(err)
		}
		city = cityName
	}

	if year == "" {
		yearName, err := ui.Ask("Year", &input.Options{
			// Read the default val from env var
			Required:  true,
			Loop:      true,
			HideOrder: true,
		})
		if err != nil {
			log.Fatal(err)
		}
		year = yearName
	}

	name, err := ui.Ask("Speaker Name", &input.Options{
		Required:  true,
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	website, err := ui.Ask("Website (optional)", &input.Options{
		HideOrder: true,
		Required:  false,
		Loop:      true,
		ValidateFunc: func(s string) error {
			if (s != "") && (helpers.ValidateField(s, "website") != true) {
				return fmt.Errorf("please enter a properly formed URL")
			}

			return nil
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	twitter, err := ui.Ask("Twitter (optional)", &input.Options{
		HideOrder: true,
		Loop:      true,
		Required:  false,
		ValidateFunc: func(s string) error {
			if (s != "") && (helpers.ValidateField(s, "twitter") != true) {
				return fmt.Errorf("please enter a properly formed twitter handle")
			}

			return nil
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	facebook, err := ui.Ask("Facebook (optional)", &input.Options{
		HideOrder: true,
		Loop:      true,
		Required:  false,
		ValidateFunc: func(s string) error {
			if (s != "") && (helpers.ValidateField(s, "facebook") != true) {
				return fmt.Errorf("please enter a properly formed Facebook profile URL")
			}

			return nil
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	linkedin, err := ui.Ask("LinkedIn (optional)", &input.Options{
		HideOrder: true, Loop: true,
		Required: false,
		ValidateFunc: func(s string) error {
			if (s != "") && (helpers.ValidateField(s, "linkedin") != true) {
				return fmt.Errorf("please enter a properly formed LinkedIn URL")
			}

			return nil
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	github, err := ui.Ask("GitHub (optional)", &input.Options{
		HideOrder: true, Loop: true,
		Required: false,
		ValidateFunc: func(s string) error {
			if (s != "") && (helpers.ValidateField(s, "github") != true) {
				return fmt.Errorf("please enter a properly formed GitHub handle")
			}

			return nil
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	gitlab, err := ui.Ask("GitLab (optional)", &input.Options{
		HideOrder: true,
		Loop:      true,
		Required:  false,
		ValidateFunc: func(s string) error {
			if (s != "") && (helpers.ValidateField(s, "gitlab") != true) {
				return fmt.Errorf("please enter a properly formed GitLab handle")
			}

			return nil
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	imageQuery, err := ui.Ask("Would you like to add the speaker image now? [Y/n]", &input.Options{
		Default:   "Y",
		HideOrder: true,
		Loop:      true,
		Required:  true,
		ValidateFunc: func(s string) error {
			if s != "Y" && s != "y" && s != "N" && s != "n" {
				return fmt.Errorf("input must be Y or n")
			}

			return nil
		},
	})

	if (imageQuery == "Y") || (imageQuery == "y") {
		imagePath = CreateSpeakerImage(helpers.NameClean(name), city, year)
		fmt.Println(imagePath)
	} else {
		imagePath = ""
	}

	fmt.Println(imageQuery)

	mySpeaker := model.Speaker{
		Name:      helpers.NameClean(name),
		Title:     name,
		Website:   website,
		Twitter:   twitter,
		Facebook:  facebook,
		Linkedin:  linkedin,
		Github:    github,
		Gitlab:    gitlab,
		ImagePath: imagePath,
	}

	NewSpeaker(mySpeaker, city, year)

	return
}

// NewSpeaker takes in a constructed Speaker type and generates the stuff
func NewSpeaker(speaker model.Speaker, city string, year string) (err error) {

	cleanName := helpers.NameClean(speaker.Name)
	t := template.New("Speaker template")

	t, err = t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	err = t.Execute(os.Stdout, speaker)
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}

	s := []string{strings.TrimSpace(cleanName), ".md"}
	f, err := os.Create(filepath.Join(helpers.EventContentPath(city, year), "speakers", strings.Join(s, "")))
	if err != nil {
		return err
	}
	defer f.Close()
	t.Execute(f, speaker)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created speaker file for", speaker.Title, "at", filepath.Join(helpers.EventContentPath(city, year), "speakers", strings.Join(s, "")))
	}
	return
}

// CreateSpeakerImage takes in a path to an image and resizes it to the proper dimensions and copies it to the destination
func CreateSpeakerImage(speaker, city, year string) (imageFile string) {
	ui := &input.UI{}
	srcPath, err := ui.Ask("Path to speaker image. Must be a PNG or JPG file.", &input.Options{
		Required:  true,
		Loop:      true,
		HideOrder: true,
		ValidateFunc: func(s string) error {
			if (s != "") && (helpers.ValidateField(s, "filepath") != true) {
				return fmt.Errorf("please enter a proper path")
			}

			if _, err := os.Stat(s); err != nil {
				return fmt.Errorf("File not found.")
			}

			return nil
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	var eventStaticPath string
	eventStaticPath, err = helpers.EventStaticPath(city, year)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(eventStaticPath)

	if err := os.MkdirAll(filepath.Join(eventStaticPath, "speakers"), 0777); err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`\.[^.]+$`)
	ext := strings.ToLower(re.FindString(srcPath))
	fmt.Println("extension is " + ext)
	switch ext {
	case ".jpg":
		s := []string{strings.TrimSpace(speaker), ".jpg"}
		destPath := filepath.Join(eventStaticPath, "speakers", strings.Join(s, ""))
		helpers.ResizeImage(srcPath, destPath, "jpg", 600, 600)
		return strings.Join(s, "")
	case ".jpeg":
		s := []string{strings.TrimSpace(speaker), ".jpg"}
		destPath := filepath.Join(eventStaticPath, "speakers", strings.Join(s, ""))
		helpers.ResizeImage(srcPath, destPath, "jpg", 600, 600)
		return strings.Join(s, "")
	case ".png":
		s := []string{strings.TrimSpace(speaker), ".png"}
		destPath := filepath.Join(eventStaticPath, "speakers", strings.Join(s, ""))
		helpers.ResizeImage(srcPath, destPath, "png", 600, 600)
		return strings.Join(s, "")
	}
	return "busted"
}
