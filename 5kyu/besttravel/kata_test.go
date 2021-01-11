package besttravel

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBestTravelKata(t *testing.T) {
	Convey("Scenario: given a list of distances ls", t, func() {
		ls := []int{50, 55, 56, 57, 58}
		Convey("Handle basic case with list 50, 55, 56, 57, 58", func() {
			So(ChooseBestSum(163, 3, ls), ShouldEqual, 163)
		})

		Convey("When list has single element 50", func() {
			ls = []int{50}
			Convey("Then it should return -1", func() {
				So(ChooseBestSum(163, 3, ls), ShouldEqual, -1)
			})
		})

		Convey("When list is 91, 74, 73, 85, 73, 81, 87", func() {
			ls = []int{91, 74, 73, 85, 73, 81, 87}
			Convey("Then it should handle several basic cases", func() {
				So(ChooseBestSum(230, 3, ls), ShouldEqual, 228)
				So(ChooseBestSum(331, 2, ls), ShouldEqual, 178)
				So(ChooseBestSum(331, 4, ls), ShouldEqual, 331)
				So(ChooseBestSum(331, 5, ls), ShouldEqual, -1)
			})
		})
	})
}
