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
	})
}
