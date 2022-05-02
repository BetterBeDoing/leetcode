package leetcode

/*Merge
先根据intervals.start进行排序
当merged[latest][end]<intervals[index][start] , 挂载新的上去
当merged[latest][end]>=intervals[index][start],可以分为两种情况
	1、当merged[latest][end]>=intervals[index][end],不更新现有的end
	2、当merged[latest][end]<intervals[index][end],更新end
*/
func Merge(intervals [][]int) [][]int {
	quickSort56(intervals, 0, len(intervals)-1)
	merged := make([][]int, 0)
	for index := range intervals {
		if index == 0 {
			merged = append(merged, intervals[index])
		}
		if merged[len(merged)-1][1] < intervals[index][0] {
			merged = append(merged, intervals[index])
		}
		if merged[len(merged)-1][1] >= intervals[index][0] {
			if merged[len(merged)-1][1] < intervals[index][1] {
				merged[len(merged)-1][1] = intervals[index][1]
			}
		}
	}
	return merged
}

func quickSort56(intervals [][]int, low int, high int) {
	if high <= low {
		return
	}
	mid := part56(intervals, low, high)
	quickSort56(intervals, low, mid-1)
	quickSort56(intervals, mid+1, high)
}

func part56(intervals [][]int, low int, high int) (mid int) {
	key := intervals[low][0]
	start := low
	end := high
	for {
		for intervals[start][0] <= key {
			start++
			if start == high {
				break
			}
		}
		for intervals[end][0] >= key {
			end--
			if end == low {
				break
			}
		}
		if start >= end {
			break
		}
		exchange(intervals, start, end)
	}
	exchange(intervals, low, end)
	return end
}

func exchange(intervals [][]int, start int, end int) {
	intervals[start], intervals[end] = intervals[end], intervals[start]
}
