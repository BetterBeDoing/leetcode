package leetcode

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestSearch(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		x := []int{4, 5, 6, 7, 0, 1, 2}
		res := Search(x, 3)
		convey.So(res, convey.ShouldEqual, 1)
	})
}
