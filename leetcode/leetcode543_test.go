package leetcode

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	convey.Convey("Boundary", t, func() {
		root := new(TreeNode)
		root.Val = 1
		root.Left = new(TreeNode)
		root.Left.Val = 2

		DiameterOfBinaryTree(root)
	})
}
