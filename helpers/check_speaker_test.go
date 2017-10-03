package helpers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckSpeaker(t *testing.T) {

	Convey("Given a speaker name", t, func() {
		Convey("When the speaker exists", func() {
			speaker := "apple-jack"
			Convey("Then the result should be true", func() {
				So(CheckSpeaker("ponyville", "2017", speaker), ShouldEqual, true)
			})
		})
		Convey("When the speaker does not exist", func() {
			speaker := "orange-jack"
			Convey("Then the response should be false", func() {
				So(CheckSpeaker("ponyville", "2017", speaker), ShouldEqual, false)
			})
		})
	})

}
