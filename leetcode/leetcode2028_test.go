package leetcode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMissingRolls(t *testing.T) {
	convey.Convey("t", t, func() {
		a := []int{3, 5, 3}
		mean := 5
		n := 3
		MissingRolls(a, mean, n)
	})
}
