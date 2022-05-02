package leetcode

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMerge(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		x := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
		m := Merge(x)
		fmt.Println(m)
	})
}
