package leetcode

func RemoveElement(nums []int, val int) int {
	var (
		ptr0 int = 0
		ptr1 int = len(nums)
		res  int = len(nums)
	)
	for ; ptr0 < ptr1; ptr0++ {
		if nums[ptr0] == val {
			res--
			nums[ptr0], nums[ptr1-1] = nums[ptr1-1], nums[ptr0]
			ptr1--
			ptr0--
		}
	}

	return res
}
