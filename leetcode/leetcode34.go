package leetcode

func SearchRange(nums []int, target int) []int {
	res := make([]int, 2)

	low := 0
	high := len(nums) - 1

	leftIdx := -1
	rightIdx := -1
	//2次二分查找
	//第一次查找nums[leftIdx]==target&&(leftIdx==0||nums[leftIdx-1]!=target)
	//第二次查找nums[rightIdx]==target&&(rightIdx==len(nums)-1||nums[rightIdx+1]!=target)
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target && (mid == 0 || nums[mid-1] != target) {
			leftIdx = mid
			break
		}
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	low = 0
	high = len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target && (mid == len(nums)-1 || nums[mid+1] != target) {
			rightIdx = mid
			break
		}
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	res[0] = leftIdx
	res[1] = rightIdx
	return res
}
