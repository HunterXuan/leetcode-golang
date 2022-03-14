package main

import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	sort.Ints(nums)

	result := make([][]int, 0)

	idx := 0
	numsLen := len(nums)
	for idx < numsLen {
		if idx > 0 && nums[idx] == nums[idx - 1] {
			idx = idx + 1
			continue
		}

		v := make([]int, len(nums))
		copy(v, nums)
		v[0], v[idx] = v[idx], v[0]

		next := permuteUnique(v[1:])
		for _, arr := range next {
			result = append(result, append([]int{nums[idx]}, arr...))
		}

		idx = idx + 1
	}

	return result
}
