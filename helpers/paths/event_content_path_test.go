package paths

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEventContentPath(t *testing.T) {

	Convey("Given a city of New York and a year of 2018", t, func() {
		city := "New York"
		year := "2018"

		testContentPath := EventContentPath(city, year)

		Convey("The response should be "+GetWebdir()+"/content/events/2018-new-york", func() {
			So(testContentPath, ShouldEqual, GetWebdir()+"/content/events/2018-new-york")
		})
	})
}
