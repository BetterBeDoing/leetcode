package leetcode

func flatten(root *TreeNode) {
	preTraverse(root)
}

func preTraverse(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = curr.Right
			curr.Left, curr.Right = nil, curr.Left
		}
		curr = curr.Right
	}
}
