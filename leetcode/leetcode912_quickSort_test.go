package leetcode

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSortArray(t *testing.T) {
	Convey("QuickSort", t, func() {
		nums := []int{3, 1}
		SortArray(nums)
		fmt.Println(nums)
	})
}
