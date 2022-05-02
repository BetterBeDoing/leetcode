package sort

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSelectSort(t *testing.T) {
	convey.Convey("in", t, func() {
		input := []int{4, 1, 5, 9, -1, 2}
		a := SelectionSort(input)
		fmt.Println(a)
	})
}
