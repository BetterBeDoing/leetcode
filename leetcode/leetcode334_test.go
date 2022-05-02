package leetcode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/*
顺便测试了一下Convey
*/
func TestIncreasingTriplet(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		x := []int{0, 1, 2, 3, 4, 5}
		res := IncreasingTriplet(x)
		convey.So(res, convey.ShouldBeTrue)
	})
}
