package leetcode

func LongestConsecutive(nums []int) int {
	return longestConsecutive(nums)
}

func longestConsecutive(nums []int) int {
	mem := make(map[int]bool, 0)
	for _, val := range nums {
		mem[val] = true
	}

	maxLen := 0
	for index := range mem {
		_, ok0 := mem[index-1]
		if !ok0 {
			tmpLen := 0
			for i := index; i < len(mem)+index; i++ {
				_, ok1 := mem[i]
				if ok1 {
					tmpLen++
				} else {
					break
				}
			}
			if tmpLen >= maxLen {
				maxLen = tmpLen
			}
		}
	}
	return maxLen
}
