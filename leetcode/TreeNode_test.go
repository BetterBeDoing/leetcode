package leetcode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBuildTree(t *testing.T) {
	convey.Convey("BuildTree", t, func() {
		nums := []int{0, 1, 2, 3, 4, 5, 6}
		BuildTree(nums)
	})
}
