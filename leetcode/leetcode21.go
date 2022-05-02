package leetcode

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//l1==nil||l2==nil
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var head *ListNode //代表已经有序的部分

	if list1.Val <= list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}
	res := head

	for {
		if list1 == nil {
			head.Next = list2
			break
		} else if list2 == nil {
			head.Next = list1
			break
		}
		if list1.Val <= list2.Val {
			head.Next = list1
			list1 = list1.Next
			head = head.Next
		} else {
			head.Next = list2
			list2 = list2.Next
			head = head.Next
		}
	}

	return res
}
