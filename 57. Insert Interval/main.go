package main

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	for _, item := range intervals {
		lastResultIdx := len(result) - 1
		if lastResultIdx >= 0 && result[lastResultIdx][1] >= item[0] {
			if item[1] > result[lastResultIdx][1] {
				result[lastResultIdx][1] = item[1]
			}
		} else {
			result = append(result, item)
		}
	}

	return result
}