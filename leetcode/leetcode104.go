package leetcode

func maxDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
		depth++
	}
	return depth
}
