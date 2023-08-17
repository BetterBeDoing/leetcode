package leetcode

func RemoveDuplicatesWithSum(nums []int) int {
	var (
		ptr0 int  = 0
		ptr1 int  = 1
		meet bool = false
	)

	for ptr1 < len(nums) {
		if nums[ptr0] == nums[ptr1] && meet {
			ptr0++
			ptr1++
			meet = true
		} else {
			ptr0++
			nums[ptr0] = nums[ptr1]
			if nums[ptr0] != nums[ptr1] {
				meet = false
			}
		}
	}

	return ptr0 + 1
}
