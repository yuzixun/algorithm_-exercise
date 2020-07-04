package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, columns := len(grid), len(grid[0])
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == '1' {
				helper(&grid, i, j, rows, columns)
				count++
			}
		}
	}
	fmt.Println(count)
	return count
}

func helper(grid *[][]byte, row, column, rows, columns int) {
	if row < 0 || column < 0 || row >= rows || column >= columns {
		return
	}
	if (*grid)[row][column] == '0' {
		return
	}
	(*grid)[row][column] = '0'

	helper(grid, row-1, column, rows, columns)
	helper(grid, row+1, column, rows, columns)
	helper(grid, row, column-1, rows, columns)
	helper(grid, row, column+1, rows, columns)
}

// func main() {
// 	numIslands([][]byte{
// 		{'1', '1', '1', '1', '0'},
// 		{'1', '1', '0', '1', '0'},
// 		{'1', '1', '0', '0', '0'},
// 		{'1', '0', '1', '0', '1'},
// 	})
// }
