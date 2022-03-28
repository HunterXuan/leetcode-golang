package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return generatePartTrees(1, n)
}

func generatePartTrees(start int, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	} else if start == end {
		return []*TreeNode{
			&TreeNode{
				Val:   start,
				Left:  nil,
				Right: nil,
			},
		}
	}

	result := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		leftNodes, rightNodes := generatePartTrees(start, i - 1), generatePartTrees(i + 1, end)
		for _, leftNode := range leftNodes {
			for _, rightNode := range rightNodes {
				result = append(result, &TreeNode{
					Val:   i,
					Left:  leftNode,
					Right: rightNode,
				})
			}
		}
	}

	return result
}