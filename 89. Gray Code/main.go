package main

func grayCode(n int) []int {
	if n <= 0 {
		return []int{0}
	}

	if n == 1 {
		return []int{0, 1}
	}

	s1 := grayCode(n - 1)
	s2 := make([]int, len(s1) * 2)
	copy(s2, s1)
	for i := 0; i < len(s1); i++ {
		s2[len(s1) + i] = 1 << (n - 1) + s1[len(s1) - 1 - i]
	}

	return s2
}
