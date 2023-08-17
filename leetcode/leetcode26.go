package leetcode

func RemoveDuplicates(nums []int) int {
	var (
		ptr0 int = 0
		ptr1 int = 1
		sum  int = 0
	)

	for ptr1 < len(nums) {
		if nums[ptr0] == nums[ptr1] && sum < 2 {
			ptr1++
			sum++
		} else {
			ptr0++
			nums[ptr0] = nums[ptr1]
			if sum >= 2 {
				sum--
			}
		}
	}

	return ptr0 + 1
}
