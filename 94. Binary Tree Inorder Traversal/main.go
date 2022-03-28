package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	if root.Left != nil {
		leftVal := inorderTraversal(root.Left)
		result = append(result, leftVal...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		rightVal := inorderTraversal(root.Right)
		result = append(result, rightVal...)
	}

	return result
}