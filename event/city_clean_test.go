package event

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCityClean(t *testing.T) {

	Convey("Given the city name of New York", t, func() {
		city := "New York"
		Convey("The result should be new-york", func() {
			So(CityClean(city), ShouldEqual, "new-york")
		})

	})
}
