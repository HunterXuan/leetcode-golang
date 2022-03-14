package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	sort.Ints(candidates)

	dfs(candidates, target, []int{}, &result)

	return result
}

func dfs(candidates []int, target int, prev []int, result *[][]int) {
	for idx, num := range candidates {
		if num == target {
			*result = append(*result, append(append([]int{}, prev...), num))
		} else if num < target {
			dfs(candidates[idx:], target - num, append(prev, num), result)
		}
	}
}