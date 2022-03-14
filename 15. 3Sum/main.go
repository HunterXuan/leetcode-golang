package main

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int

	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 {
			break
		}

		target := -nums[i]
		l, r := i+1, n-1
		for l < r {
			tmp := nums[l] + nums[r]
			if tmp == target {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l = l + 1
				r = r - 1
				for l < r && nums[l] == nums[l-1] {
					l = l + 1
				}
				for r > l && nums[r] == nums[r+1] {
					r = r - 1
				}
			} else if tmp > target {
				r = r - 1
			} else {
				l = l + 1
			}
		}
	}

	return result
}