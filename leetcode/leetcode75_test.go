package leetcode

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSortColors(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		x := []int{0, 2, 1, 2}
		SortColors(x)
		fmt.Println(x)
	})
}
