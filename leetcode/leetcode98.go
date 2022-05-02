package leetcode

import "math"

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MaxInt, math.MinInt)
}

func isValid(root *TreeNode, max int, min int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	lf := isValid(root.Left, root.Val, min)
	rf := isValid(root.Right, max, root.Val)
	return lf && rf
}

/*func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}*/
