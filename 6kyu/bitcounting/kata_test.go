package bitcounting

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBitCountingKata(t *testing.T) {
	Convey("Given a number", t, func() {
		numbers := []uint{0, 4, 7, 9, 10}

		Convey("When CountBits() is called", func() {
			results := make([]int, len(numbers))
			for i, n := range numbers {
				results[i] = CountBits(n)
			}
			Convey("Then get number of 1s in that number's binary representation", func() {
				expects := []int{0, 1, 3, 2, 2}
				for i, r := range results {
					Convey(fmt.Sprintf("CountBits(%d) should be equal to %d\n", numbers[i], expects[i]), func() {
						So(r, ShouldEqual, expects[i])
					})
				}
			})
		})
	})
}
