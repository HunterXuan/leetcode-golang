package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	row := []*TreeNode{root}
	for len(row) > 0 {
		tmpRow := make([]*TreeNode, 0)
		tmpResult := make([]int, 0)
		for _, node := range row {
			tmpResult = append(tmpResult, node.Val)

			if node.Left != nil {
				tmpRow = append(tmpRow, node.Left)
			}

			if node.Right != nil {
				tmpRow = append(tmpRow, node.Right)
			}
		}

		result = append(result, tmpResult)
		row = tmpRow
	}

	return result
}