package stringtocamelcase

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStringToCamelCaseKata(t *testing.T) {
	Convey("Given a string", t, func() {
		s1 := "the-stealth-warrior"
		s2 := "The_Stealth_Warrior"
		Convey("Solution should return camel cased string", func() {
			r1 := ToCamelCase(s1)
			r2 := ToCamelCase(s2)

			e1 := "theStealthWarrior"
			e2 := "TheStealthWarrior"

			So(r1, ShouldEqual, e1)
			So(r2, ShouldEqual, e2)
		})
	})
}
