package main

func rotate(matrix [][]int)  {
	n := len(matrix)
	mid := (n - 1) >> 1
	for r := 0; r <= mid; r++ {
		for c := r; c < n-r-1; c++ {
			matrix[c][n-r-1], matrix[n-r-1][n-c-1], matrix[n-c-1][r], matrix[r][c] = matrix[r][c], matrix[c][n-r-1], matrix[n-r-1][n-c-1], matrix[n-c-1][r]
		}
	}
}