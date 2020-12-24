package stringendswith

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKata(t *testing.T) {
	Convey("Given a string 's'", t, func() {
		s := "abc"

		Convey("When checked with string containing 'bc'", func() {
			ending := "bc"
			result := solution(s, ending)

			Convey("The result should be true", func() {
				So(result, ShouldBeTrue)
			})
		})

		Convey("When checked with string containing 'd'", func() {
			ending := "d"
			result := solution(s, ending)

			Convey("The result should be false", func() {
				So(result, ShouldBeFalse)
			})
		})
	})
}
