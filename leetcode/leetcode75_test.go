package leetcode

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestSortColors(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		x := []int{0, 2, 1, 2}
		SortColors(x)
		fmt.Println(x)
	})
}
