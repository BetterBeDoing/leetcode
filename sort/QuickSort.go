package sort

func quickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}
