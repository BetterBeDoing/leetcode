package leetcode

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestQuickSortByH(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		x := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
		QuickSortByH(x, 0, len(x)-1)
		fmt.Println(x)
	})
}
