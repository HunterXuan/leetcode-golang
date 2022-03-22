package main

func combine(n int, k int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	if k == 1 {
		tmp := make([][]int, n)
		for i := 0; i < n; i++ {
			tmp[i] = []int{i + 1}
		}
		return tmp
	}

	if n == k {
		tmp := make([]int, n)
		for i := 0; i < n; i++ {
			tmp[i] = i + 1
		}
		return [][]int{tmp}
	}

	result := make([][]int, 0)

	tmp := combine(n - 1, k - 1)
	for _, item := range tmp {
		result = append(result, append(item, n))
	}

	result = append(result, combine(n - 1, k)...)

	return result
}