package findodd

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindOddKata(t *testing.T) {
	Convey("Given an array of integers", t, func() {
		arr := [...][]int{
			{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5},
			{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5},
			{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5},
			{10},
			{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1},
			{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10},
		}

		Convey("When FindOdd() is called with that array", func() {
			results := make([]int, len(arr))
			for i, a := range arr {
				results[i] = FindOdd(a)
			}

			Convey("It should return the integer that appears odd number of times", func() {
				expects := [...]int{5, -1, 5, 10, 10, 1}

				for i, a := range arr {
					Convey(fmt.Sprintf("FindOdd(%v) should return %d\n", a, expects[i]), func() {
						So(results[i], ShouldEqual, expects[i])
					})
				}
			})
		})
	})
}
