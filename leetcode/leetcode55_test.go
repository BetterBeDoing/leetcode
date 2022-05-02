package leetcode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCanJump(t *testing.T) {
	convey.Convey("In", t, func() {
		CanJump([]int{2, 0})
	})
}
