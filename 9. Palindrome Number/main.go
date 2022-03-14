package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x >= 0 && x <= 9 {
		return true
	}

	var arr []int
	for x > 0 {
		arr = append(arr, x % 10)
		x = x / 10
	}

	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		if arr[i] != arr[arrLen - i - 1] {
			return false
		}
	}

	return true
}