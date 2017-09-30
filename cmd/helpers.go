package cmd

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func cityClean(city string) (cityClean string) {
	cityClean = strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)
	return
}

func eventDataPath(webdir, city, year string) (eventDataPath string) { // TODO: Add argument for webdir path
	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10), ".yml"}
	eventDataPath = filepath.Join(webdir, "data", "events", strings.Join(s, ""))
	// eventDataPath = strings.Join(s, "")
	// eventDataPath = webdir
	return eventDataPath
}

func eventContentPath(webdir, city, year string) (eventContentPath string) { // TODO: Add argument for webdir path
	s := []string{strings.TrimSpace(year), "-", strings.Replace(strings.TrimSpace(strings.ToLower(city)), " ", "-", 10)}
	eventContentPath = filepath.Join(webdir, "content", "events", strings.Join(s, ""))
	// eventContentPath = webdir
	return eventContentPath
}

func validateField(input, field string) bool {
	switch field {
	case "city":
		if strings.Count(input, "") > 100 {
			return false
		}
		return true
	case "year":
		if strings.Count(input, "") != 5 {
			return false
		} else if s, err := strconv.ParseInt(input, 10, 32); err == nil {
			if s < 2016 || s > 2030 {
				return false
			}
			return true

		}
	case "twitter":
		if strings.ContainsAny(input, " ") {
			return false
		}
		return true
	}
	return true // TODO: Make this return an error if no field was matched
}

// checkEvent takes in two arguments, the city and the year, and returns true if the city  exists.
func checkEvent(city, year string) bool {
	if _, err := os.Stat(eventDataPath(webdir, city, year)); err == nil {
		return true
	}
	return false

}

func fieldMap() (fieldMap map[string]string) {
	tempMap := make(map[string]string)
	tempMap["EventTwitter"] = "Twitter"
	tempMap["GaTrackingID"] = "Google Analytics Tracking ID"
	tempMap["Startdate"] = "Start Date"
	tempMap["Enddate"] = "End Date"
	tempMap["CfpDateStart"] = "CFP Start Date"
	tempMap["CfpDateEnd"] = "CFP End Date"
	tempMap["CfpDateAnnounce"] = "CFP Announcement Date"
	tempMap["CfpOpen"] = "CFP Link"
	tempMap["RegistrationDateStart"] = "Registation Start Date"
	tempMap["RegistrationDateEnd"] = "Registration End Date"
	tempMap["RegistrationLink"] = "Registration Link"
	tempMap["Coordinates"] = "Coordinates"
	tempMap["Location"] = "Location"
	tempMap["LocationAddress"] = "Location Address"

	return tempMap
}
