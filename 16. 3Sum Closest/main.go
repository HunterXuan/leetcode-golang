package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	result := nums[0] + nums[1] + nums[2]
	gap := math.MaxInt

	sort.Ints(nums)

	numsLen := len(nums)
	for i := 0; i < len(nums); i++ {
		l, r := i + 1, numsLen - 1
		for l < r {
			tmp := nums[i] + nums[l] + nums[r]
			diff := tmp - target
			if diff < 0 {
				diff = -diff
			}

			if diff < gap {
				result = tmp
				gap = diff
			}

			if tmp == target {
				return target
			} else if tmp < target {
				l = l + 1
			} else {
				r = r - 1
			}
		}
	}

	return result
}