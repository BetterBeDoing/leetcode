package leetcode

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestWordBreak(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		/*
			"catsandog"
			["cats","dog","sand","and","cat"]
		*/
		x := []string{"a", "dog", "sand", "and", "cat"}
		s := "a"
		res := WordBreak(s, x)
		convey.So(res, convey.ShouldBeTrue)
	})
}
