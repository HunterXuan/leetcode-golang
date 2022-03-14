package main

func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for idx, _ := range matrix {
		matrix[idx] = make([]int, n)
	}

	matrix[m - 1][n - 1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			tmp := 0

			if i < m - 1 {
				tmp = tmp + matrix[i + 1][j]
			}

			if j < n - 1 {
				tmp = tmp + matrix[i][j + 1]
			}

			if tmp > 0 {
				matrix[i][j] = tmp
			}
		}
	}

	return matrix[0][0]
}