package main

type ListNode struct {
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	left, right := &ListNode{}, &ListNode{}
	pointer, leftPointer, rightPointer := head, left, right
	for pointer != nil {
		nextPointer := pointer.Next
		pointer.Next = &ListNode{}
		if pointer.Val < x {
			leftPointer.Next = pointer
			leftPointer = pointer
		} else {
			rightPointer.Next = pointer
			rightPointer = pointer
		}

		pointer = nextPointer
	}

	rightPointer.Next = nil
	leftPointer.Next = right.Next

	return left.Next
}