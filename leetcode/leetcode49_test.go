package leetcode

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGroupAnagrams(t *testing.T) {
	convey.Convey("In", t, func() {
		strs := []string{"ddddddddddg", "dgggggggggg"}
		res := [][]string{{"dgggggggggg"}, {"ddddddddddg"}}
		x := GroupAnagrams(strs)
		convey.So(x, convey.ShouldResemble, res)
	})
}
