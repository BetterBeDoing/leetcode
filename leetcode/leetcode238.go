package leetcode

func productExceptSelf(nums []int) []int {
	var (
		L   = make([]int, len(nums))
		R   = make([]int, len(nums))
		ans = make([]int, len(nums))
	)

	for i := 0; i <= len(nums)-1; i++ {
		if i == 0 {
			L[i] = 1
			R[len(nums)-1-i] = 1
		} else {
			L[i] = L[i-1] * nums[i-1]
			R[len(nums)-1-i] = R[len(nums)-i] * nums[len(nums)-i]
		}
	}

	for i := 0; i <= len(nums)-1; i++ {
		ans[i] = L[i] * R[i]
	}

	return ans
}
