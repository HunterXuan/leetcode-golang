package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode = nil
	return isBSTHelper(&pre, root)
}

func isBSTHelper(pre **TreeNode, root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBSTHelper(pre, root.Left) {
		return false
	}

	if *pre != nil && (*pre).Val >= root.Val {
		return false
	}

	*pre = root
	return isBSTHelper(pre, root.Right)
}