package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	for _, item := range intervals {
		resultLastIdx := len(result) - 1
		if resultLastIdx >= 0 && item[0] <= result[resultLastIdx][1] {
			if item[1] > result[resultLastIdx][1] {
				result[resultLastIdx][1] = item[1]
			}
		} else {
			result = append(result, item)
		}
	}

	return result
}