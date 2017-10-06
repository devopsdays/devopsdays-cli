package helpers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidateField(t *testing.T) {
	// if v := ValidateField("devopsdays", "twitter"); v != true {
	// 	t.Error("Valid twitter did not pass validation test in ValidateField")
	// }
	// if v := ValidateField("devops days", "twitter"); v == true {
	// 	t.Error("Invalid twitter passed validation test in ValidateField")
	// }
	Convey("Given a validation request for a city", t, func() {
		Convey("When the city name is 'Chicago'", func() {
			Convey("Then the response should be true", func() {
				So(ValidateField("Chicago", "city"), ShouldEqual, true)
			})
		})
		Convey("When the city name is 3yl0RmG1wU8q5TeDPKZEsNU3E54nyYf5MNhGhzqcxhoLJkeckXCa1saWCPM24YhwIteGEUjLW8S715WkoDvt3vFsMaVeYXCUZWNL", func() {
			Convey("Then the response should be false", func() {
				So(ValidateField("3yl0RmG1wU8q5TeDPKZEsNU3E54nyYf5MNhGhzqcxhoLJkeckXCa1saWCPM24YhwIteGEUjLW8S715WkoDvt3vFsMaVeYXCUZWNL", "city"), ShouldEqual, false)
			})
		})
	})

	Convey("Given a validation request for a year", t, func() {
		Convey("When the year is 2017", func() {
			year := "2017"
			Convey("Then the response should be true", func() {
				So(ValidateField(year, "year"), ShouldEqual, true)
			})
		})
		Convey("When the year is 19008", func() {
			year := "19008"
			Convey("Then the response should be false", func() {
				So(ValidateField(year, "year"), ShouldEqual, false)
			})
		})
		Convey("When the year is 2031", func() {
			year := "2031"
			Convey("Then the response should be false", func() {
				So(ValidateField(year, "year"), ShouldEqual, false)
			})
		})
	})

	Convey("Given a validation request for a twitter account", t, func() {
		Convey("When the account name is 'devopsdays'", func() {
			twitter := "devopsdays"
			Convey("Then the response should be true", func() {
				So(ValidateField(twitter, "twitter"), ShouldEqual, true)
			})
		})
		Convey("When the account name is devops days", func() {
			twitter := "devops days"
			Convey("Then the response should be false", func() {
				So(ValidateField(twitter, "twitter"), ShouldEqual, false)
			})
		})
		Convey("When the account name is @devopsdays", func() {
			twitter := "@devopsdays"
			Convey("Then the response should be true", func() {
				So(ValidateField(twitter, "twitter"), ShouldEqual, true)
			})
		})
	})

	Convey("Given a validation request for a Facebook setting", t, func() {
		Convey("When the setting is 'https://www.facebook.com/matt.stratton'", func() {
			facebook := "https://www.facebook.com/matt.stratton"
			Convey("Then the response should be true", func() {
				So(ValidateField(facebook, "facebook"), ShouldEqual, true)
			})
		})
		Convey("When the setting is 'matt.stratton'", func() {
			facebook := "matt.stratton"
			Convey("Then the response should be false", func() {
				So(ValidateField(facebook, "facebook"), ShouldEqual, false)
			})
		})
	})
	Convey("Given a validation request for a LinkedIn setting", t, func() {
		Convey("When the setting is 'https://www.linkedin.com/in/mattstratton/'", func() {
			linkedin := "https://www.linkedin.com/in/mattstratton/"
			Convey("Then the response should be true", func() {
				So(ValidateField(linkedin, "linkedin"), ShouldEqual, true)
			})
		})
		Convey("When the setting is 'matt.stratton'", func() {
			linkedin := "matt.stratton"
			Convey("Then the response should be false", func() {
				So(ValidateField(linkedin, "linkedin"), ShouldEqual, false)
			})
		})
	})
	Convey("Given a validation request for a GitHub setting", t, func() {
		Convey("When the setting is 'mattstratton'", func() {
			github := "mattstratton"
			Convey("Then the response should be true", func() {
				So(ValidateField(github, "github"), ShouldEqual, true)
			})
		})
		Convey("When the setting is 'https://github.com/mattstratton'", func() {
			github := "https://github.com/mattstratton"
			Convey("Then the response should be false", func() {
				So(ValidateField(github, "github"), ShouldEqual, false)
			})
		})
	})
	Convey("Given a validation request for a GitLab setting", t, func() {
		Convey("When the setting is 'mattstratton'", func() {
			gitlab := "mattstratton"
			Convey("Then the response should be true", func() {
				So(ValidateField(gitlab, "gitlab"), ShouldEqual, true)
			})
		})
		Convey("When the setting is 'https://gitlab.com/matt.stratton'", func() {
			gitlab := "https://gitlab.com/matt.stratton"
			Convey("Then the response should be false", func() {
				So(ValidateField(gitlab, "gitlab"), ShouldEqual, false)
			})
		})
	})
	Convey("Given a validation request for a filepath", t, func() {
		Convey("When the file path is '/home/mattstratton/foo'", func() {
			filepath := "/home/mattstratton/foo"
			Convey("Then the response should be true", func() {
				So(ValidateField(filepath, "filepath"), ShouldEqual, true)
			})
		})
		Convey("When the file path is 'c:/home/mattstratton'", func() {
			filepath := "c:/home/mattstratton'"
			Convey("Then the response should be false", func() {
				So(ValidateField(filepath, "filepath"), ShouldEqual, false)
			})
		})
		Convey("When the file path is 'c:\\home\\mattstratton'", func() {
			filepath := "c:\\home\\mattstratton'"
			Convey("Then the response should be true", func() {
				So(ValidateField(filepath, "filepath"), ShouldEqual, true)
			})
		})
		Convey("When the file path is 'mything:stuff'", func() {
			filepath := "mything:stuff"
			Convey("Then the response should be false", func() {
				So(ValidateField(filepath, "filepath"), ShouldEqual, false)
			})
		})
	})
}
