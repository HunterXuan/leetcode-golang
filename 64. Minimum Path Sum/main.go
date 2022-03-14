package main

func minPathSum(grid [][]int) int {
	rowCount, colCount := len(grid), len(grid[0])
	resultGrid := make([][]int, rowCount)
	for idx, _ := range resultGrid {
		resultGrid[idx] = make([]int, colCount)
	}

	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			if i > 0 && j > 0 {
				if resultGrid[i-1][j] > resultGrid[i][j-1] {
					resultGrid[i][j] = grid[i][j] + resultGrid[i][j - 1]
				} else {
					resultGrid[i][j] = grid[i][j] + resultGrid[i - 1][j]
				}
			} else if i > 0 {
				resultGrid[i][j] = grid[i][j] + resultGrid[i - 1][j]
			} else if j > 0 {
				resultGrid[i][j] = grid[i][j] + resultGrid[i][j - 1]
			} else {
				resultGrid[i][j] = grid[i][j]
			}
		}
	}

	return resultGrid[rowCount - 1][colCount - 1]
}