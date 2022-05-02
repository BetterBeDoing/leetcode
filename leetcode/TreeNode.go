package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(nums []int) (root *TreeNode) {
	l := len(nums)
	root = new(TreeNode)
	queue := make([]*TreeNode, l)

	if l == 0 {
		return nil
	} else {
		start := 0
		end := start
		for start < l {
			if end == 0 {
				root.Val = nums[start]
				queue[0] = root
				end++
			} else {
				if end < l {
					left := new(TreeNode)
					left.Val = nums[end]
					queue[start].Left = left
					queue[end] = left
					end++
				}
				if end < l {
					right := new(TreeNode)
					right.Val = nums[end]
					queue[start].Right = right
					queue[end] = right
					end++
				}
				start++
			}
		}
	}
	return root
}
