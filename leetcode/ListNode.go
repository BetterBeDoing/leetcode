package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode() *ListNode {
	return new(ListNode)
}

func (l *ListNode) AddList(input []int) {
	end := l
	for _, val := range input {
		if end.Next != nil {
			end = end.Next
		}
		end.Val = val
		next := new(ListNode)
		end.Next = next
	}
	end.Next = nil
}

func (l ListNode) Traverse() {
	for {
		fmt.Println(l.Val)
		if l.Next != nil {
			l = *l.Next
		} else {
			break
		}
	}
}

func (l *ListNode) TransToNums() (res []int) {
	for {
		res = append(res, l.Val)
		if l.Next != nil {
			l = l.Next
		} else {
			break
		}
	}
	return res
}
