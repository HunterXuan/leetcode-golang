package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	var result [][]int

	sort.Ints(nums)

	numsLen := len(nums)
	for i := 0; i < numsLen - 3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		tmp := threeSum(nums[i+1:numsLen], target - nums[i])
		for _, item := range tmp {
			result = append(result, append(item, nums[i]))
		}
	}

	return result
}

func threeSum(nums []int, target int) [][]int {
	var result [][]int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i + 1, len(nums) - 1
		for l < r {
			tmpSum := nums[i] + nums[l] + nums[r]
			if  tmpSum == target {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l = l + 1
				r = r - 1
				for l < r && nums[l] == nums[l-1] {
					l = l + 1
				}
				for r > l && nums[r] == nums[r+1] {
					r = r - 1
				}
			} else if tmpSum < target {
				l = l + 1
			} else {
				r = r - 1
			}
		}
	}

	return result
}