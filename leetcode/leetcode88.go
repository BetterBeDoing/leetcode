package leetcode

// TIme: O(ptr3) = O(m+n-1) = O(m+n)
// Spac: O(1)
/***
其实这题很简单，因为nums1的末尾是有空位的，我们只需要从末尾开始执行就可以将多余的数填入进去。
***/
func merge(nums1 []int, m int, nums2 []int, n int) {
	var ptr1, ptr2, ptr3 int = m - 1, n - 1, m + n - 1
	for ; ptr1 >= 0 && ptr2 >= 0; ptr3-- {
		if nums2[ptr2] >= nums1[ptr1] {
			nums1[ptr3] = nums2[ptr2]
			ptr2--
		} else {
			nums1[ptr3] = nums1[ptr1]
			ptr1--
		}
	}

	if ptr2 >= 0 {
		copy(nums1[:ptr3+1], nums2[:ptr2+1])
	}
}
