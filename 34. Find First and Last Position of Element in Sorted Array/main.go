package main

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}

	numsLen := len(nums)
	if numsLen <= 0 || nums[0] > target || nums[numsLen - 1] < target {
		return result
	}

	l, r, mid := 0, numsLen - 1, 0
	for l < r {
		mid = (l + r) >> 1
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		} else {
			break
		}
	}

	mid = (l + r) >> 1
	if nums[mid] != target {
		return result
	}

	result = []int{mid, mid}
	for l = mid; l >= 0; l-- {
		if nums[l] == target {
			result[0] = l
		} else {
			break
		}
	}

	for r = mid; r < numsLen; r++ {
		if nums[r] == target {
			result[1] = r
		} else {
			break
		}
	}

	return result
}