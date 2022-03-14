package main

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	pointer := &ListNode{}
	header := &ListNode{Next: pointer}
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pointer.Val = list1.Val
			list1 = list1.Next
		} else {
			pointer.Val = list2.Val
			list2 = list2.Next
		}

		pointer.Next = &ListNode{}
		pointer = pointer.Next
	}

	if list1 != nil {
		pointer.Val = list1.Val
		pointer.Next = list1.Next
	}

	if list2 != nil {
		pointer.Val = list2.Val
		pointer.Next = list2.Next
	}

	return header.Next
}