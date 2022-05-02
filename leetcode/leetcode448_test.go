package leetcode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	convey.Convey("tets", t, func() {

		nums := []int{2, 2}
		res := FindDisappearedNumbers(nums)
		res0 := []int{1}
		convey.So(res, convey.ShouldResemble, res0)
	})
}
