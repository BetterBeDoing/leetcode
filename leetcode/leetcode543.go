package leetcode

func DiameterOfBinaryTree(root *TreeNode) int {
	max := 0
	depth(root, &max)
	return max
}

func depth(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left, max)
	right := depth(root.Right, max)
	dep := left + right
	if dep > *max {
		*max = dep
	}
	if left > right {
		return left + 1
	}
	return right + 1
}
