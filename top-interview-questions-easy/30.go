package main

func isValidSudoku(board [][]byte) bool {
	rows := make([]int, 9)
	columns := make([]int, 9)
	square := make([]int, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			value := board[i][j] - '0' - 1
			cur := 1 << value
			sIndex := (i / 3) + (j/3)*3
			if cur&(rows[i]|columns[j]|square[sIndex]) == cur {
				return false
			}

			rows[i] |= cur
			columns[j] |= cur
			square[sIndex] |= cur
		}
	}

	return true
}
