package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	Convey("1", t, func() {
		s := "cbaebabacd"
		p := "abc"
		res := FindAnagrams(s, p)
		So(res, ShouldResemble, []int{0, 6})
	})
}
