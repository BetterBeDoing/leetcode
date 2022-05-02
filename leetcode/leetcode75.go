package leetcode

//mark0标记0的左边界
//mark2标记2的右边界
func SortColors(nums []int) {
	mark0 := 0
	mark2 := len(nums) - 1
	left := mark0
	for left <= mark2 {
		if nums[left] == 0 {
			nums[left], nums[mark0] = nums[mark0], nums[left]
			mark0++
		}
		if nums[left] == 2 {
			nums[left], nums[mark2] = nums[mark2], nums[left]
			mark2--
		}
		if nums[left] == 0 && left != mark0-1 || nums[left] == 2 {
			continue
		} else {
			left++
		}
	}
}
