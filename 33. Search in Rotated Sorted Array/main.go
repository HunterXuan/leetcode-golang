package main

func search(nums []int, target int) int {
	numsLen := len(nums)

	l, r := 0, numsLen - 1
	for l < r {
		n := (l + r) >> 1
		if n == l || n == r {
			if nums[l] == target {
				return l
			} else if nums[r] == target {
				return r
			} else {
				return -1
			}
		}

		if nums[n] >= nums[l] {
			if nums[l] <= target && target <= nums[n] {
				r = n
			} else {
				l = n
			}
		} else {
			if nums[n] <= target && target <= nums[r] {
				l = n
			} else {
				r = n
			}
		}
	}

	if nums[r] == target {
		return r
	}

	return -1
}