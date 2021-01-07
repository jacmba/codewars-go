package encryptthis

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncryptThisKata(t *testing.T) {
	Convey("Given a string", t, func() {
		s1 := ""
		s2 := "A wise old owl lived in an oak"
		s3 := "The more he saw the less he spoke"
		s4 := "The less he spoke the more he heard"
		s5 := "Why can we not all be like that wise old bird"
		s6 := "Thank you Piotr for all your help"

		Convey("Should encrypt each word", func() {
			r1 := EncryptThis(s1)
			r2 := EncryptThis(s2)
			r3 := EncryptThis(s3)
			r4 := EncryptThis(s4)
			r5 := EncryptThis(s5)
			r6 := EncryptThis(s6)

			e1 := ""
			e2 := "65 119esi 111dl 111lw 108dvei 105n 97n 111ka"
			e3 := "84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp"
			e4 := "84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare"
			e5 := "87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri"
			e6 := "84kanh 121uo 80roti 102ro 97ll 121ruo 104ple"

			So(r1, ShouldEqual, e1)
			So(r2, ShouldEqual, e2)
			So(r3, ShouldEqual, e3)
			So(r4, ShouldEqual, e4)
			So(r5, ShouldEqual, e5)
			So(r6, ShouldEqual, e6)
		})
	})
}
