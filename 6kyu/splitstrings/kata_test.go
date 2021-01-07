package splitstrings

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSplitStringsKata(t *testing.T) {
	Convey("Given a string", t, func() {
		s1 := "abc"
		s2 := "abcdef"

		Convey("Should return array of char pairs padding with underscore", func() {
			r1 := Solution(s1)
			r2 := Solution(s2)

			So(r1, ShouldResemble, []string{"ab", "c_"})
			So(r2, ShouldResemble, []string{"ab", "cd", "ef"})
		})
	})
}
