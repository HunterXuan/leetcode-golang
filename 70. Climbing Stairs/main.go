package main

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	result, a, b := 0, 1, 1
	for i := 2; i <= n; i++ {
		result = a + b
		a, b = result, a
	}

	return result
}