package leetcode

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 != nil && root2 == nil {
		return root1
	} else if root1 == nil && root2 != nil {
		return root2
	} else {
		root1.Val += root2.Val
	}

	if root1.Left != nil && root2.Left != nil {
		mergeTrees(root1.Left, root2.Left)
	} else {
		if root2.Left != nil {
			root1.Left = root2.Left
		}
	}
	if root1.Right != nil && root2.Right != nil {
		mergeTrees(root1.Right, root2.Right)
	} else {
		if root2.Right != nil {
			root1.Right = root2.Right
		}
	}
	return root1
}
