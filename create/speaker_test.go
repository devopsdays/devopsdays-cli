package create

import (
	"testing"

	"github.com/devopsdays/devopsdays-cli/model"
)

func TestNewSpeaker(t *testing.T) {
	t.Log("Testing NewSpeaker function...")
	mySpeaker := model.Speaker{
		Name:      "matt-stratton",
		Title:     "Matt Stratton",
		Website:   "http://www.mattstratton.com",
		Twitter:   "mattstratton",
		Facebook:  "https://www.facebook.com/matt.stratton",
		Linkedin:  "https://www.linkedin.com/in/mattstratton/",
		Github:    "mattstratton",
		Gitlab:    "mattstratton",
		ImagePath: "matt-stratton.jpg",
	}
	err := NewSpeaker(mySpeaker, "ponyville", "2017")
	if err != nil {
		t.Errorf("Received error on NewSpeaker")
	}
}
