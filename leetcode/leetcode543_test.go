package leetcode

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
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
