package leetcode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	two := head
	for head != nil {
		head = head.Next
		if two != nil {
			if two.Next != nil {
				two = two.Next.Next
			} else {
				two = two.Next
			}
		}
		if head == two && head != nil {
			return true
		}
	}
	return false
}
