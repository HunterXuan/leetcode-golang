package main

type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	oriHead, p1, p2 := head, head, head
	for tmp := k; tmp > 0; {
		p2 = p2.Next
		tmp = tmp - 1
		if tmp > 0 && p2.Next == nil {
			listLen := k - tmp + 1
			tmp = k % listLen
			p2 = head
		}
	}

	if p1 == p2 {
		return head
	}

	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	result := p1.Next
	p1.Next = nil
	p2.Next = oriHead

	return result
}