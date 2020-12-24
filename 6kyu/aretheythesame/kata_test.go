package aretheythesame

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAreTheyTheSameKata(t *testing.T) {
	Convey("Given 2 arrays of integers", t, func() {
		a := []int{121, 144, 19, 161, 19, 144, 19, 11}
		b := []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}

		Convey("When array B contains squares of A elements", func() {
			Convey("Then the arrays should be same", func() {
				result := Comp(a, b)
				So(result, ShouldBeTrue)
			})
		})

		Convey("When array B contains other random elements", func() {
			b = []int{132, 14641, 20736, 361, 25921, 361, 20736, 361}
			Convey("Then the arrays should not be same", func() {
				result := Comp(a, b)
				So(result, ShouldBeFalse)
			})
		})
	})
}
