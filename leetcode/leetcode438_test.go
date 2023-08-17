package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindAnagrams(t *testing.T) {
	Convey("1", t, func() {
		s := "cbaebabacd"
		p := "abc"
		res := FindAnagrams(s, p)
		So(res, ShouldResemble, []int{0, 6})
	})
}
