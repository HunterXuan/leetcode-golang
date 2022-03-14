package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	result := make([][]int, 0)

	dfs(candidates, target, []int{}, &result)

	return result
}

func dfs(candidates []int, target int, prev []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, prev...))
	}

	if target < 0 {
		return
	}

	for i := 0; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}

		dfs(candidates[i+1:], target - candidates[i], append(prev, candidates[i]), result)
	}
}