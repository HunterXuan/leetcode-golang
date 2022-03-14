package main

func spiralOrder(matrix [][]int) []int {
	rowCount, colCount := len(matrix), len(matrix[0])
	result := make([]int, rowCount * colCount)

	visitMap := make([][]bool, rowCount)
	for i := range visitMap {
		visitMap[i] = make([]bool, colCount)
	}

	idx, row, col, direction := 0, 0, 0, "right"
	for idx < rowCount * colCount {
		result[idx] = matrix[row][col]
		visitMap[row][col] = true
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

	return result
}