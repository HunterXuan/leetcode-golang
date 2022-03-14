package main

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	for n > 0 {
		p2 = p2.Next
		n = n - 1
	}

	if p2 == nil {
		if p1.Next != nil {
			return p1.Next
		}
		return nil
	}

	for p2.Next != nil {
		p2 = p2.Next
		p1 = p1.Next
	}

	p1.Next = p1.Next.Next

	return head
}