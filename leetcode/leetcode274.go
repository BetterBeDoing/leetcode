package leetcode

// use bucket sort
func hIndex(citations []int) int {
	var (
		bucket = make([]int, 1000)
	)
	for _, v := range citations {
		bucket[v]++
	}

	for i := len(bucket) - 1; i > 0; i-- {
		if bucket[i] >= i {
			return i
		}
		bucket[i-1] += bucket[i]
	}

	return 0
}

//3,0,6,1,5
//0,1,3,5,6
//1 1 1 1 1
//5 4 3 2 1

//1 3 1
//1 2 3
//2 2 1
