package helpers

import (
	"strconv"
	"strings"
)

// ValidateField performs field validation for common types
func ValidateField(input, field string) bool {
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
