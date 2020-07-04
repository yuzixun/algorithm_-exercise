package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows, columns := len(grid), len(grid[0])
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == '1' {
				numIslandsHelper(&grid, rows, columns, i, j)
				count++
			}
		}
	}

	return count
}

func numIslandsHelper(grid *[][]byte, rows, columns, row, column int) {
	if row < 0 || column < 0 || row > rows-1 || column > columns-1 || (*grid)[row][column] == '0' {
		return
	}
	(*grid)[row][column] = '0'
	numIslandsHelper(grid, rows, columns, row+1, column)
	numIslandsHelper(grid, rows, columns, row-1, column)
	numIslandsHelper(grid, rows, columns, row, column+1)
	numIslandsHelper(grid, rows, columns, row, column-1)
}

// func main() {
// 	numIslands([][]byte{
// 		[]byte{'1', '1', '1', '1', '0'},
// 		[]byte{'1', '1', '0', '1', '0'},
// 		[]byte{'1', '1', '0', '0', '0'},
// 		[]byte{'0', '0', '0', '0', '0'},
// 	})

// 	numIslands([][]byte{
// 		[]byte{'1', '1', '0', '0', '0'},
// 		[]byte{'1', '1', '0', '0', '0'},
// 		[]byte{'0', '0', '1', '0', '0'},
// 		[]byte{'0', '0', '0', '1', '1'},
// 	})

// }
