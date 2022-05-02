package leetcode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/*func TestJudgeReverse(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		x := "asa"
		y := "assa"
		res := JudgeReverse(x)
		res1 := JudgeReverse(y)
		convey.So(res, convey.ShouldEqual, 1)
		convey.So(res1, convey.ShouldEqual, 1)

	})
}*/

func TestCountSubstrings(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		x := "abc"
		y := "assa"
		res := CountSubstrings(x)
		res1 := CountSubstrings(y)
		convey.So(res1, convey.ShouldEqual, 6)
		convey.So(res, convey.ShouldEqual, 3)

	})
}
