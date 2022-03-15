package main

func setZeroes(matrix [][]int)  {
	position := make([][]int, 0)
	rowCount, colCount := len(matrix), len(matrix[0])
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			if matrix[i][j] == 0 {
				position = append(position, []int{i, j})
			}
		}
	}

	for _, item := range position {
		row, col := item[0], item[1]
		for i := 0; i < colCount; i++ {
			matrix[row][i] = 0
		}
		for i := 0; i < rowCount; i++ {
			matrix[i][col] = 0
		}
	}
}