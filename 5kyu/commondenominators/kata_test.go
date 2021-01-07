package commondenominators

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCommonDenominators(t *testing.T) {
	Convey("Given a list of fractions as pair arrays", t, func() {
		l1 := [][]int{{1, 2}, {1, 3}, {1, 4}}
		l2 := [][]int{{69, 130}, {87, 1310}, {30, 40}}

		Convey("Solution should return common factor fractions as string", func() {
			r1 := ConvertFracts(l1)
			r2 := ConvertFracts(l2)

			e1 := "(6,12)(4,12)(3,12)"
			e2 := "(18078,34060)(2262,34060)(25545,34060)"

			So(r1, ShouldEqual, e1)
			So(r2, ShouldEqual, e2)
		})
	})
}
