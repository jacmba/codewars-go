package multiples3or5

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMultiples3Or5Kata(t *testing.T) {
	Convey("Given a number", t, func() {
		n := 10
		Convey("Solution should return the sum of multiples of 3 and 5 below that number", func() {
			r := Multiple3And5(n)
			So(r, ShouldEqual, 23)
		})
	})
}
