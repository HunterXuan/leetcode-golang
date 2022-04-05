package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		nodesLen := len(nodes)
		tmpNodes := make([]*TreeNode, 0)
		for idx, node := range nodes {
			if !isEqualNode(node, nodes[nodesLen - idx - 1]) {
				return false
			}

			if node != nil {
				tmpNodes = append(tmpNodes, node.Left, node.Right)
			}
		}
		nodes = tmpNodes
	}

	return true
}

func isEqualNode(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val
}