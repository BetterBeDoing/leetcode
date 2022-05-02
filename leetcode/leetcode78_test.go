package leetcode

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSubsets(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		x := []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 0}
		m := Subsets(x)
		fmt.Println(m)
	})
}
