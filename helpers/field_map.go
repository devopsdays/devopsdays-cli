package helpers

// FieldMap returns a mapping of field titles to pretty ones?
func FieldMap() (fieldMap map[string]string) {
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
