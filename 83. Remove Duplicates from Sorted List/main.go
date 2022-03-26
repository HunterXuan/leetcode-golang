package main

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	pointer := head
	for pointer != nil {
		if pointer.Next != nil && pointer.Val == pointer.Next.Val {
			pointer.Next = pointer.Next.Next
		} else {
			pointer = pointer.Next
		}
	}

	return head
}