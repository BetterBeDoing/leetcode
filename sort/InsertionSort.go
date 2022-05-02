package sort

func InsertionSort(nums []int) []int {
	sorted := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[sorted] {
			tmp := nums[i]
			x := sorted
			for {
				if x > -1 && tmp < nums[x] {
					next := x + 1
					nums[next] = nums[x]
					x--
				} else {
					if x == -1 {
						nums[0] = tmp
						break
					}
					nums[x] = tmp
					break
				}
			}
		}
		sorted++
	}
	return nums
}
