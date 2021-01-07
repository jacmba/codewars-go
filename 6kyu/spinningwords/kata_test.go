package spinningwords

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpinningWordsKata(t *testing.T) {
	Convey("Given a string", t, func() {
		s1 := "Hey fellow warriors"
		s2 := "This is a test"
		s3 := "This is another test"
		Convey("Solution should return words bigger than 5 chars with reversed characters", func() {
			r1 := SpinWords(s1)
			r2 := SpinWords(s2)
			r3 := SpinWords(s3)

			e1 := "Hey wollef sroirraw"
			e2 := "This is a test"
			e3 := "This is rehtona test"

			So(r1, ShouldEqual, e1)
			So(r2, ShouldEqual, e2)
			So(r3, ShouldEqual, e3)
		})
	})
}
