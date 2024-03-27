package leetcode

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestLongestCommonPrefix(t *testing.T) {
	convey.Convey("test", t, func() {
		res := LongestCommonPrefix([]string{"flower", "flow", "flight"})
		convey.So(res, convey.ShouldEqual, "fl")
	})
}
