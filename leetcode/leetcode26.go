package leetcode

func RemoveDuplicates(nums []int) int {
	var (
		ptr0 int = 0
		ptr1 int = 1
	)

	for ptr1 < len(nums) {
		if nums[ptr0] == nums[ptr1] {
			ptr1++
			sum++
		} else {
			ptr0++
			nums[ptr0] = nums[ptr1]
		}
	}

	return ptr0 + 1
}
