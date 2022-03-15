package main

func searchMatrix(matrix [][]int, target int) bool {
	rowCount, colCount := len(matrix), len(matrix[0])
	if target < matrix[0][0] || target > matrix[rowCount - 1][colCount - 1] {
		return false
	}

	l, r := 0, rowCount * colCount - 1
	for l < r {
		mid := (l + r) >> 1
		if matrix[mid / colCount][mid % colCount] == target {
			return true
		} else if matrix[mid / colCount][mid % colCount] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return matrix[r / colCount][r % colCount] == target
}