package leetcode

func MidRoot(root *TreeNode) []int {
	var res []int
	var midRoot func(root *TreeNode)
	midRoot = func(root *TreeNode) {
		if root == nil {
			return
		}
		midRoot(root.Left)
		res = append(res, root.Val)
		midRoot(root.Right)
	}
	midRoot(root)
	return res
}

func PreRoot(root *TreeNode) []int {
	var res []int
	var preRoot func(root *TreeNode)
	preRoot = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preRoot(root.Left)
		preRoot(root.Right)
	}
	preRoot(root)
	return res
}

func PostRoot(root *TreeNode) []int {
	var res []int
	var postRoot func(root *TreeNode)
	postRoot = func(root *TreeNode) {
		if root == nil {
			return
		}
		postRoot(root.Left)
		postRoot(root.Right)
		res = append(res, root.Val)
	}
	postRoot(root)
	return res
}
