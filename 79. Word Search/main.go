package main

func exist(board [][]byte, word string) bool {
	rowCount, colCount := len(board), len(board[0])

	var spread func(i int, j int, w string) bool
	spread = func(i int, j int, w string) bool {
		if len(w) == 0 {
			return true
		}

		original := board[i][j]
		board[i][j] = '-'
		flag := false
		for _, item := range [][]int{{i-1, j},{i+1, j}, {i, j-1}, {i, j+1}} {
			x, y := item[0], item[1]
			if x >= 0 && x < rowCount && y >= 0 && y < colCount && w[0] == board[x][y] && spread(x, y, w[1:]) {
				flag = true
				break
			}
		}

		board[i][j] = original
		return flag
	}

	for r := 0; r < rowCount; r++ {
		for c := 0; c < colCount; c++ {
			if word[0] == board[r][c] && spread(r, c, word[1:]) {
				return true
			}
		}
	}

	return false
}