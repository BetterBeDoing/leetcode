package leetcode

func Rotate(nums []int, k int) {
	if len(nums) <= 1 || k == 0 {
		return
	}
	k %= len(nums)
	//cannot use a new memory
	reverse_189(nums)
	reverse_189(nums[0:k])
	reverse_189(nums[k:len(nums)])
}

func reverse_189(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}
