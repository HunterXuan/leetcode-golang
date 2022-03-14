package main

func maxArea(height []int) int {
	l, r := 0, len(height) - 1
	maxVal := (r - l) * min(height[l], height[r])
	for l < r {
		if height[r] > height[l] {
			l = l + 1
		} else {
			r = r - 1
		}

		maxVal = max(maxVal, (r - l) * min(height[l], height[r]))
	}

	return maxVal
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}