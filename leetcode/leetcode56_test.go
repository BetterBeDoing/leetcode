package leetcode

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestMerge(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		x := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
		m := Merge(x)
		fmt.Println(m)
	})
}
