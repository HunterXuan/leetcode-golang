package main

func isValidSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		occurMap := make(map[byte]bool)
		for col := 0; col < 9; col++ {
			tmp := board[row][col]
			if tmp == '.' {
				continue
			} else if occurMap[tmp] {
				return false
			} else {
				occurMap[tmp] = true
			}
		}
	}

	for col := 0; col < 9; col++ {
		occurMap := make(map[byte]bool)
		for row := 0; row < 9; row++ {
			tmp := board[row][col]
			if tmp == '.' {
				continue
			} else if occurMap[tmp] {
				return false
			} else {
				occurMap[tmp] = true
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			occurMap := make(map[byte]bool)
			for k := 0; k < 9; k++ {
				row, col := k / 3 + i * 3, k - (k / 3) * 3 + j * 3
				tmp := board[row][col]
				if tmp == '.' {
					continue
				} else if occurMap[tmp] {
					return false
				} else {
					occurMap[tmp] = true
				}
			}
		}
	}

	return true
}