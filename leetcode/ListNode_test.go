package leetcode

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestNewNode(t *testing.T) {
	convey.Convey("Test", t, func() {
		l := NewNode()
		x1 := []int{1, 3, 4, 5, 9}
		l.AddList(x1)
		l.Traverse()
	})
}
