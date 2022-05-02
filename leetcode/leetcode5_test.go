package leetcode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	convey.Convey("test", t, func() {
		res0 := LongestPalindrome("abbba")
		convey.So(res0, convey.ShouldEqual, "abbba")
		res1 := LongestPalindrome("bb")
		convey.So(res1, convey.ShouldEqual, "bb")
	})
}
