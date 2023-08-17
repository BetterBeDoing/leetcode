package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveDuplicatesWithSum(t *testing.T) {
	Convey("in", t, func() {
		input := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
		res := RemoveDuplicatesWithSum(input)
		target_res := []int{1, 1, 2, 2, 3, 3}
		Convey(ShouldResemble, res, target_res)
	})
}
