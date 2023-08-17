package leetcode

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestBoolLongestPalindrome(t *testing.T) {
	convey.Convey("in", t, func() {
		x1 := "abbba"
		res := BoolLongestPalindrome(x1)
		convey.So(res, convey.ShouldEqual, "abbba")
	})
}
