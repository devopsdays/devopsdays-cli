package helpers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckSpeaker(t *testing.T) {

	Convey("Given a speaker name", t, func() {
		speaker := "apple-jack"
		Convey("When the speaker exists", func() {
			result := CheckSpeaker("ponyville", "2017", speaker)
			Convey("Then the result should be true", func() {
				So(result, ShouldEqual, true)
			})
		})
	})

	Convey("Given a speaker name", t, func() {
		speaker := "orange-jack"
		Convey("When the speaker does not exist", func() {
			result := CheckSpeaker("ponyville", "2017", speaker)
			Convey("Then the result should be false", func() {
				So(result, ShouldEqual, false)
			})
		})
	})
}
