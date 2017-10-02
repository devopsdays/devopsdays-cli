package helpers

import (
	"testing"
)

func TestValidateField(t *testing.T) {
	t.Log("Testing ValidateField() function")
	if v := ValidateField("Chicago", "city"); v != true {
		t.Error("Valid city did not pass validation test in ValidateField")
	}
	if v := ValidateField("3yl0RmG1wU8q5TeDPKZEsNU3E54nyYf5MNhGhzqcxhoLJkeckXCa1saWCPM24YhwIteGEUjLW8S715WkoDvt3vFsMaVeYXCUZWNL", "city"); v == true {
		t.Error("Invalid city passed validation test in ValidateField")
	}
	if v := ValidateField("2017", "year"); v != true {
		t.Error("Valid year did not pass validation test in ValidateField")
	}
	if v := ValidateField("19008", "year"); v == true {
		t.Error("Invalid year passed validation test in ValidateField")
	}
	if v := ValidateField("2011", "year"); v == true {
		t.Error("Invalid year passed validation test in ValidateField")
	}
	if v := ValidateField("2031", "year"); v == true {
		t.Error("Invalid year passed validation test in ValidateField")
	}
	if v := ValidateField("devopsdays", "twitter"); v != true {
		t.Error("Valid twitter did not pass validation test in ValidateField")
	}
	if v := ValidateField("devops days", "twitter"); v == true {
		t.Error("Invalid twitter passed validation test in ValidateField")
	}
}
