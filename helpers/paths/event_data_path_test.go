package paths

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEventDataPath(t *testing.T) {

	Convey("Given a city of New York and a year of 2018", t, func() {
		city := "New York"
		year := "2018"

		testContentPath := EventDataPath(GetWebdir(), city, year)

		Convey("The response should be "+GetWebdir()+"/data/events/2018-new-york.yml", func() {
			So(testContentPath, ShouldEqual, GetWebdir()+"/data/events/2018-new-york.yml")
		})
	})
}
