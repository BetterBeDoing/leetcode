package leetcode

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestSearchRange(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		//x := []int{5, 5, 7, 7, 8, 8, 10, 10}
		x := []int{5, 5}
		res := SearchRange(x, 5)
		convey.So(res, convey.ShouldResemble, []int{0, 1})
	})
}
