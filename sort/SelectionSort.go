package sort

func SelectionSort(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		min := nums[i]
		minPos := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				minPos = j
			}
		}
		tmp := nums[i]
		nums[i] = min
		nums[minPos] = tmp
	}
	return nums
}
