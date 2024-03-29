package main

func search(nums []int, target int) bool {
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return true
		}

		if nums[l] == nums[mid] {
			l = l + 1
			continue
		} else if nums[r] == nums[mid] {
			r = r - 1
			continue
		}

		if nums[l] < nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false
}