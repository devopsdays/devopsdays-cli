package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/devopsdays/devopsdays-cli/helpers"
	"github.com/devopsdays/devopsdays-cli/helpers/paths"
	"github.com/devopsdays/devopsdays-cli/model"
	"gopkg.in/yaml.v2"
)

var city string
var year string

func organizerFieldMap() (fieldMap map[string]string) {
	tempMap := make(map[string]string)
	tempMap["Name"] = "Organizer Name"
	tempMap["Twitter"] = "Twitter name (without @ symbol)"
	tempMap["Employer"] = "Optional Employer Name"
	tempMap["Github"] = "GitHub Username"
	tempMap["Facebook"] = "Facebook URL"
	tempMap["Linkedin"] = "Linkedin URL"
	tempMap["Website"] = "URL to personal website"
	tempMap["Image"] = "image name"
	tempMap["Bio"] = "Bio - markdown allowed"

	return tempMap
}

func eventFields() []string {
	fields := make([]string, 14)
	fields[0] = "EventTwitter"
	fields[1] = "GaTrackingID"
	fields[2] = "Startdate"
	fields[3] = "Enddate"
	fields[4] = "CfpDateStart"
	fields[5] = "CfpDateEnd"
	fields[6] = "CfpDateAnnounce"
	fields[7] = "CfpOpen"
	fields[8] = "RegistrationDateStart"
	fields[9] = "RegistrationDateEnd"
	fields[10] = "RegistrationLink"
	fields[11] = "Coordinates"
	fields[12] = "Location"
	fields[13] = "LocationAddress"

	return fields
}

func organizerFields() []string {
	fields := make([]string, 9)
	fields[0] = "Name"
	fields[1] = "Twitter"
	fields[2] = "Employer"
	fields[3] = "Github"
	fields[4] = "Facebook"
	fields[5] = "Linkedin"
	fields[6] = "Website"
	fields[7] = "Image"
	fields[8] = "Bio"

	return fields
}

func editEvent(event model.Event) (err error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you want to [1] edit the value of a field, [2] add an organizer, or [3] add a sponsor?")
	c, _ := reader.ReadString('\n')
	c = strings.TrimSpace(c)
	switch c {
	case "1":
		s := eventFields()
		myField := makeMenu(s)
		fmt.Println("The value of this field is: ", returnField(event, myField))
		fmt.Println("Would you like to change it?")
		c, _ := reader.ReadString('\n')
		c = strings.TrimSpace(c)
		if c == "y" {
			fmt.Println("What would you like to change it to?")
			c, _ := reader.ReadString('\n')
			c = strings.TrimSpace(c)
			editField(event, myField, c)
		}
	case "2":
		fmt.Println("The list of organizers is:")
		m := make(map[int]string)
		for o, value := range event.TeamMembers {
			s := reflect.ValueOf(&value).Elem()
			typeOfT := s.Type()
			for i := 0; i < s.NumField(); i++ {
				f := s.Field(i)
				if typeOfT.Field(i).Name == "Name" {
					fmt.Print("[", (o + 1), "] ", f.Interface(), "\n")
					n := f.String()
					m[o+1] = n
				}
			}
		}
		fmt.Println("Who would you like to see more about?")
		c, _ := reader.ReadString('\n')
		c = strings.TrimSpace(c)
		c2, _ := strconv.Atoi(c)
		o := m[c2]
		for _, value := range event.TeamMembers {
			s := reflect.ValueOf(&value).Elem()
			r := reflect.Indirect(s).FieldByName("Name")
			r2 := r.String()
			if r2 == o {
				typeOfT := s.Type()
				for i := 0; i < s.NumField(); i++ {
					f := s.Field(i)
					fmt.Print(typeOfT.Field(i).Name, ": ")
					fmt.Print(f.Interface(), "\n")
					updateOrganizer(event, o, "Twitter", "Dude")
				}
			}
		}

	case "3":
		// fmt.Println("Adding sponsors is not yet supported.")
		fmt.Print(event.NavElements)
		fmt.Print("The length is", len(event.NavElements))
		myOrg := organizerStruct("matt", "mattstratton", "", "mattstratton", "fb", "li", "website", "img", "stuff and junk")
		fmt.Print(myOrg)
		spew.Dump(myOrg)
		spew.Dump(event.TeamMembers)
		for _, value := range event.TeamMembers {
			s := reflect.ValueOf(&value).Elem()
			fmt.Println("-------------")
			fmt.Println(value)
			fmt.Println("-------------")
			fmt.Println(s.Field(0))
		}

	default:
		fmt.Println("This is the default.")
	}

	// Note: This is commented out for now, but we want to use this functionality somewhere
	// event.Name = "mugsyville"
	// fmt.Print(event.Name)
	// y, err := yaml.Marshal(&event)
	// ioutil.WriteFile((helpers.EventDataPath(webdir, city, year)), y, 0755)

	return

}

func eventStruct(city, year string) (event model.Event) {
	// var event Event
	yamlFile, err := ioutil.ReadFile(paths.EventDataPath(webdir, city, year))
	err = yaml.Unmarshal(yamlFile, &event)
	if err != nil {
		panic(err)
	}
	return event
}

func organizerStruct(name, twitter, employer, github, facebook, linkedin, website, image, bio string) (organizer model.Organizer) {
	o := model.Organizer{Name: name, Twitter: twitter, Employer: employer, Github: github, Facebook: facebook, Linkedin: linkedin, Website: website, Image: image, Bio: bio}

	return o

}

// TODO: This should actually return the key to change; rather than just create the menu
func makeMenu(items []string) (field string) {
	fmt.Println("Which field would you like to modify?")
	myMap := helpers.FieldMap()
	menu := "\n"
	for i, v := range items {
		menu += "["
		menu += strconv.Itoa(i + 1)
		menu += "] "
		menu += myMap[v]
		menu += "\n"
	}
	fmt.Println(menu)
	reader := bufio.NewReader(os.Stdin)
	var c, _ = reader.ReadString('\n')
	c = strings.TrimSpace(c)
	c2, _ := strconv.Atoi(c)
	field = items[(c2 - 1)]

	return field
}

func returnField(event model.Event, field string) (name string) {
	r := reflect.ValueOf(event)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func editField(event model.Event, field, value string) {
	// r := reflect.ValueOf(event)
	// f := reflect.Indirect(r).FieldByName(field)
	reflect.ValueOf(&event).Elem().FieldByName(field).SetString(value)
	y, _ := yaml.Marshal(&event)
	ioutil.WriteFile((paths.EventDataPath(webdir, event.City, event.Year)), y, 0755)
	return
}

func updateOrganizer(event model.Event, name, field, value string) {
	for _, loopvalue := range event.TeamMembers {
		s := reflect.ValueOf(&loopvalue).Elem()
		r := reflect.Indirect(s).FieldByName(field)
		if (s.Field(0)).String() == name {
			r.SetString(value)
			fmt.Println(r)
			spew.Dump(event.TeamMembers)
			y, _ := yaml.Marshal(&event)
			ioutil.WriteFile((paths.EventDataPath(webdir, event.City, event.Year)), y, 0755)
		}
	}
}

func editOrganizer(event model.Event, organizer, field, value string) {
	for _, value := range event.TeamMembers {
		s := reflect.ValueOf(&value).Elem()
		fmt.Print(s)
		typeOfT := s.Type()
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fmt.Print("key: ", typeOfT.Field(i).Name, "\n")
			fmt.Print("value: ", f.Interface(), "\n")
		}
	}
}
