package main

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	pointer := &ListNode{}
	result := pointer
	for l1 != nil || l2 != nil || flag > 0 {
		pointer.Next = &ListNode{}
		pointer = pointer.Next

		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		tmp := val1 + val2 + flag
		if tmp >= 10 {
			tmp = tmp - 10
			flag = 1
		} else {
			flag = 0
		}

		pointer.Val = tmp
	}

	return result.Next
}
