package sort

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	convey.Convey("in", t, func() {
		input := []int{4, 1, 5, 9, -1, 2}
		a := BubbleSort(input)
		fmt.Println(a)
	})
}
