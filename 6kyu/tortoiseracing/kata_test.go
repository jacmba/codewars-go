package tortoiseracing

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTortoiseRacingKata(t *testing.T) {
	Convey("Scenario: tortoise racing should handle basic cases", t, func() {
		So(Race(720, 850, 70), ShouldResemble, [3]int{0, 32, 18})
		So(Race(820, 81, 550), ShouldResemble, [3]int{-1, -1, -1})
		So(Race(80, 91, 37), ShouldResemble, [3]int{3, 21, 49})
	})
}
