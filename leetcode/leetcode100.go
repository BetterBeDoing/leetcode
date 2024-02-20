package leetcode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) && p.Val == q.Val
	} else {
		return false
	}
}
