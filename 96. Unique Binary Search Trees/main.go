package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func numTrees(n int) int {
	dp := make([]int, n + 1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] = dp[i] + dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}
