package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rowCount, colCount := len(obstacleGrid), len(obstacleGrid[0])
	pathCountGrid := make([][]int, rowCount)
	for idx, _ := range pathCountGrid {
		pathCountGrid[idx] = make([]int, colCount)
	}

	pathCountGrid[rowCount - 1][colCount - 1] = 1
	for i := rowCount - 1; i >= 0; i-- {
		for j := colCount - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 0 {
				tmp := 0

				if i < rowCount - 1 && obstacleGrid[i + 1][j] == 0 {
					tmp = tmp + pathCountGrid[i + 1][j]
				}

				if j < colCount - 1 && obstacleGrid[i][j + 1] == 0 {
					tmp = tmp + pathCountGrid[i][j + 1]
				}

				if tmp > 0 {
					pathCountGrid[i][j] = tmp
				}
			} else {
				pathCountGrid[i][j] = 0
			}
		}
	}

	return pathCountGrid[0][0]
}