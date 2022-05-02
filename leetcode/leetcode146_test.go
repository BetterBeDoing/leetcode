package leetcode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestForTest(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		x := Constructor(2)
		x.Put(2, 1)
		x.Put(1, 1)
		x.Put(2, 3)
		x.Put(4, 1)
		x.Get(1)
		x.Get(2)
	})
}
