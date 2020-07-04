package main

func numRookCaptures(board [][]byte) int {
	count := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'R' {
				count = count + handle(board, i, j, 0, 1) + handle(board, i, j, 0, -1) + handle(board, i, j, -1, 0) + handle(board, i, j, 1, 0)
			}
		}
	}

	return count
}

func handle(board [][]byte, row, column, dx, dy int) int {
	for row >= 0 && len(board) > row && column >= 0 && len(board[0]) > column && board[row][column] != 'B' {
		if board[row][column] == 'p' {
			return 1
		}
		row += dx
		column += dy
	}
	return 0
}

func main() {

}
