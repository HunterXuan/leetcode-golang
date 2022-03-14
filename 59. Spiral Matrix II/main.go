package main

func generateMatrix(n int) [][]int {
	rowCount, colCount := n, n

	matrix := make([][]int, rowCount)
	for idx, _ := range matrix {
		matrix[idx] = make([]int, colCount)
	}

	visitMap := make([][]bool, rowCount)
	for idx, _ := range matrix {
		visitMap[idx] = make([]bool, colCount)
	}

	idx, row, col, direction := 0, 0, 0, "right"
	for idx < rowCount * colCount {
		visitMap[row][col] = true
		matrix[row][col] = idx + 1
		idx = idx + 1

		if direction == "right" {
			if col < colCount - 1 && !visitMap[row][col + 1] {
				col = col + 1
			} else {
				row = row + 1
				direction = "down"
			}
		} else if direction == "down" {
			if row < rowCount - 1 && !visitMap[row + 1][col] {
				row = row + 1
			} else {
				col = col - 1
				direction = "left"
			}
		} else if direction == "left" {
			if col > 0 && !visitMap[row][col - 1] {
				col = col - 1
			} else {
				row = row - 1
				direction = "up"
			}
		} else if direction == "up" {
			if row > 0 && !visitMap[row - 1][col] {
				row = row - 1
			} else {
				col = col + 1
				direction = "right"
			}
		}
	}

	return matrix
}