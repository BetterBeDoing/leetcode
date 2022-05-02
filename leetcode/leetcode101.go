package leetcode

func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

/*func check(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val && check(l.Left, r.Right) && check(l.Right, r.Left)
}*/

func check(l *TreeNode, r *TreeNode) bool {
	memL := make([]*TreeNode, 0)
	memR := make([]*TreeNode, 0)

	memL = append(memL, l)
	memR = append(memR, r)

	for len(memL) > 0 && len(memR) > 0 {
		startL := memL[0]
		startR := memR[0]
		if startL == nil && startR == nil {
			continue
		}
		if startL == nil || startR == nil {
			return false
		}
		if startL.Val != startR.Val {
			return false
		}
		memL = append(memL, startL.Left)
		memR = append(memR, startR.Right)
		memL = append(memL, startL.Right)
		memR = append(memR, startR.Left)

		memL = memL[1:]
		memR = memR[1:]
	}
	if len(memL) == 0 && len(memR) == 0 {
		return false
	}
	return true
}
