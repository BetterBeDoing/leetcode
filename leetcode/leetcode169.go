package leetcode

func MajorityElement(nums []int) int {
	var (
		mem map[int]int
		max int
		l   int
	)

	l = len(nums)
	max = nums[0]
	mem = make(map[int]int, len(nums))

	for _, v := range nums {
		sum, ok := mem[v]
		if ok {
			mem[v] = sum + 1
			if mem[v] > l/2 {
				max = v
			}
		} else {
			mem[v] = 1
		}
	}

	return max
}
