package duplicateencoder

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDuplicateEncoderKata(t *testing.T) {
	Convey("Given a string", t, func() {
		s1 := "din"
		s2 := "recede"
		s3 := "Success"
		s4 := "(( @"

		Convey("Solution should return string of ( or ) chars based on original chars occurrences", func() {
			r1 := DuplicateEncode(s1)
			r2 := DuplicateEncode(s2)
			r3 := DuplicateEncode(s3)
			r4 := DuplicateEncode(s4)

			e1 := "((("
			e2 := "()()()"
			e3 := ")())())"
			e4 := "))(("

			So(r1, ShouldEqual, e1)
			So(r2, ShouldEqual, e2)
			So(r3, ShouldEqual, e3)
			So(r4, ShouldEqual, e4)
		})
	})
}
