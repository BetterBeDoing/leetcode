package leetcode

func FindDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	res := make([]int, 0)
	for i, v := range nums {
		if v <= n {
			res = append(res, i+1)
		}
	}
	return res
}
