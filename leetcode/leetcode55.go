package leetcode

//应当选择动态规划
func CanJump(nums []int) bool {
	max := -1
	for idx := range nums {
		if idx <= max {
			nums[idx] += idx
			if nums[idx] > max {
				max = nums[idx]
			}
		}
	}
	if max > len(nums)-1 {
		return true
	}
	return false
}
