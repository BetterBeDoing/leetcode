package leetcode

import "math/rand"

func SortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, low int, high int) {
	if low >= high {
		return
	}
	pos := randomPartition(nums, low, high)
	//分两路归并排序
	quickSort(nums, low, pos-1)
	quickSort(nums, pos+1, high)
}

func randomPartition(nums []int, p, r int) int {
	i := rand.Intn(r-p) + p //找到一个随机的位置
	// fmt.Println(i)
	nums[i], nums[r] = nums[r], nums[i]
	return betterPartition(nums, p, r)
}

func betterPartition(nums []int, p, r int) int {
	x, i := nums[r], p-1
	for j := p; j < r; j++ {
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

func Divide(nums []int, low int, high int) int {
	i := low
	j := high
	for {
		i++
		for ; nums[i] <= nums[low] && i <= high; i++ {
			if i == high {
				break
			}
		}
		for ; nums[j] > nums[low] && j >= low; j-- {
			if j == low {
				break
			}
		}
		if j <= i {
			break
		}
		change(nums, i, j)
	}
	change(nums, j, low)
	return j
}

func change(nums []int, i int, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
