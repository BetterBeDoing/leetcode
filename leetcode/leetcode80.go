package leetcode

func RemoveDuplicatesWithSum(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	idx := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[idx-2] {
			nums[idx] = nums[i]
			idx += 1

		}
	}
	return idx
}
