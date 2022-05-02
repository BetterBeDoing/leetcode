package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPermute(t *testing.T) {
	Convey("1", t, func() {
		input := []int{1, 2, 3}
		ans := Permute(input)
		output := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
		So(ans, ShouldResemble, output)
	})
}
