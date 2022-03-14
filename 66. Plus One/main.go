package main

func plusOne(digits []int) []int {
	left := 0
	numsLen := len(digits)
	digits[numsLen - 1] = digits[numsLen - 1] + 1
	for idx := numsLen - 1; idx >= 0; idx-- {
		tmp := digits[idx] + left
		if tmp >= 10 {
			digits[idx] = tmp - 10
			left = 1
		} else {
			digits[idx] = tmp
			left = 0
		}
	}

	if left == 0 {
		return digits
	}

	return append([]int{left}, digits...)
}