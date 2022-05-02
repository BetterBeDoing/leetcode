package leetcode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBoolLongestPalindrome(t *testing.T) {
	convey.Convey("in", t, func() {
		x1 := "abbba"
		res := BoolLongestPalindrome(x1)
		convey.So(res, convey.ShouldEqual, "abbba")
	})
}
