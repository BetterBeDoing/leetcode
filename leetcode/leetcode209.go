package leetcode

func minSubArrayLen(target int, nums []int) int {
	var sum, left, right, minLen int
	minLen = len(nums) + 1
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}
