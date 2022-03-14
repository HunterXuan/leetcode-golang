package main

func searchInsert(nums []int, target int) int {
	numsLen := len(nums)
	l, r, mid := 0, numsLen-1, 0
	for l < r {
		mid = (l + r) >> 1
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		} else {
			return mid
		}
	}

	if nums[l] >= target {
		return l
	}

	return l+1
}