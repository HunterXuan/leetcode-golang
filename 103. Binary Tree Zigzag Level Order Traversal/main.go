package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	row := []*TreeNode{root}
	l2r := true
	for len(row) > 0 {
		tmpRow := make([]*TreeNode, 0)
		tmpResult := make([]int, 0)
		if l2r {
			for idx := len(row) - 1; idx >= 0; idx-- {
				node := row[idx]
				tmpResult = append(tmpResult, node.Val)

				if node.Left != nil {
					tmpRow = append(tmpRow, node.Left)
				}

				if node.Right != nil {
					tmpRow = append(tmpRow, node.Right)
				}
			}
		} else {
			for idx := len(row) - 1; idx >= 0; idx-- {
				node := row[idx]
				tmpResult = append(tmpResult, node.Val)

				if node.Right != nil {
					tmpRow = append(tmpRow, node.Right)
				}

				if node.Left != nil {
					tmpRow = append(tmpRow, node.Left)
				}
			}
		}

		result = append(result, tmpResult)
		row = tmpRow
		l2r = !l2r
	}

	return result
}