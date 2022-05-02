package leetcode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	convey.Convey("Test", t, func() {
		l1 := NewNode()
		l2 := NewNode()
		x1 := []int{1, 7, 13, 24, 25}
		x2 := []int{1, 2, 4, 8, 15, 22}
		l1.AddList(x1)
		l2.AddList(x2)
		res0 := MergeTwoLists(l1, l2)
		nums0 := res0.TransToNums()
		convey.So(nums0, convey.ShouldResemble, []int{1, 1, 2, 4, 7, 8, 13, 15, 22, 24, 25})
	})
}
