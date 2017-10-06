package names

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNameClean(t *testing.T) {

	Convey("Given a name of George Bluth", t, func() {
		name := "George Bluth"

		Convey("The response should be George Bluth", func() {
			So(NameClean(name), ShouldEqual, "george-bluth")
		})
	})
}
