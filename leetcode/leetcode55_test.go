package leetcode

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestCanJump(t *testing.T) {
	convey.Convey("In", t, func() {
		CanJump([]int{2, 0})
	})
}
