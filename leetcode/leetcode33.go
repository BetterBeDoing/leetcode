package leetcode

//Search 采用折半查找的方法
/*
当nums[left]>nums[right] -> [left:right]是中间含有折半的
不然则按照正常的处理
*/
func Search(nums []int, target int) int {
	ln := len(nums)
	l, r := 0, ln-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		// l, mid, r 一条线
		if nums[mid] >= nums[l] && nums[mid] <= nums[r] {
			if target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] >= nums[l] { // l, mid在左， r在右
			if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else { // l在左，mid, r在右
			if target < nums[mid] || target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
