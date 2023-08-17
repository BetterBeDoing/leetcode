package leetcode

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestLongestConsecutive(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		x := []int{0, 1, 2, 3, 4, 5}
		res := LongestConsecutive(x)
		convey.So(res, convey.ShouldEqual, 6)
	})
}
