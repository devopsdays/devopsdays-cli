package speaker

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckSpeaker(t *testing.T) {

	Convey("Given a speaker name", t, func() {
		Convey("When the speaker exists", func() {
			speaker := "nathen-harvey"
			Convey("Then the result should be true", func() {
				So(CheckSpeaker("chicago", "2019", speaker), ShouldEqual, true)
			})
		})
		Convey("When the speaker does not exist", func() {
			speaker := "hannah-foxwell"
			Convey("Then the response should be false", func() {
				So(CheckSpeaker("chicago", "2019", speaker), ShouldEqual, false)
			})
		})
	})

}
