package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	prevHead := &ListNode{
		Next: head,
	}

	pointer := prevHead
	for i := 0; i < left - 1; i++ {
		pointer = pointer.Next
	}

	leftHead := pointer
	pointer = pointer.Next
	for i := 0; i < right - left; i++ {
		n1 := leftHead.Next
		n2 := pointer.Next
		n3 := pointer.Next.Next

		leftHead.Next = n2
		n2.Next = n1
		pointer.Next = n3
	}

	return prevHead.Next
}