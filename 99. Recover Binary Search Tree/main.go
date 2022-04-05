package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode)  {
	var first, second, prevNode *TreeNode
	inorder(root, &first, &second, &prevNode)
	first.Val, second.Val = second.Val, first.Val
}

func inorder(root *TreeNode, first, second, prevNode **TreeNode) {
	if root == nil { return }
	inorder(root.Left, first, second, prevNode)
	if *prevNode != nil && (*prevNode).Val >= root.Val {
		if *first == nil {
			*first = *prevNode
		}
		*second = root
	}
	*prevNode = root
	inorder(root.Right, first, second, prevNode)
}