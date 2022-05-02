package leetcode

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestExist(t *testing.T) {
	Convey("in", t, func() {
		input := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
		target := "SEE"
		x := Exist(input, target)
		fmt.Println(x)
	})
}
