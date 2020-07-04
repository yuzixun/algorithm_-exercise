package main

func gameOfLife(board [][]int) {
	xx := len(board)
	if xx == 0 {
		return
	}
	yy := len(board[0])
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	cache := make([][]int, xx)

	for i := 0; i < xx; i++ {
		cache[i] = make([]int, yy)
	}

	x, y, count := 0, 0, 0
	for i := 0; i < xx; i++ {
		for j := 0; j < yy; j++ {
			count = 0
			for k := 0; k < 8; k++ {
				x, y = i+dx[k], j+dy[k]
				if x < 0 || x >= xx || y < 0 || y >= yy {
					continue
				}

				count += board[x][y]
			}

			switch {
			case count < 2:
				cache[i][j] = 0
			case count == 2:
				if board[i][j] == 1 {
					cache[i][j] = 1
				}
			case count == 3:
				cache[i][j] = 1
			case count > 3:
				cache[i][j] = 0
			}
		}
	}

	for i := 0; i < xx; i++ {
		for j := 0; j < yy; j++ {
			board[i][j] = cache[i][j]
		}
	}
	// fmt.Println(board)
}

func main() {
	gameOfLife(
		[][]int{
			[]int{0, 1, 0},
			[]int{0, 0, 1},
			[]int{1, 1, 1},
			[]int{0, 0, 0}})
	gameOfLife([][]int{})
}
