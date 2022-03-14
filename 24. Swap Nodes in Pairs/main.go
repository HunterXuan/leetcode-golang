package main

type ListNode struct {
	Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var p0, p1, p2, p3 *ListNode
	p0 = &ListNode{Next: head}
	result := p0
	for p0.Next != nil && p0.Next.Next != nil {
		p1, p2, p3 = p0.Next, p0.Next.Next, p0.Next.Next.Next

		p0.Next = p2
		p2.Next = p1
		p1.Next = p3

		p0 = p1
	}

	return result.Next
}