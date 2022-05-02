package leetcode

func convertBST(root *TreeNode) *TreeNode {
	dfs(root, 0)
	return root
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	r := dfs(root.Right, sum)
	sum = r + root.Val
	root.Val = sum
	l := dfs(root.Left, sum)
	return root.Val + l
}
