package main

func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	columns := len(grid[0])
	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == '1' {
				count++
				clear(grid, rows, columns, i, j)
			}
		}
	}
	// fmt.Println("count is ", count)
	return count
}

func clear(grid [][]byte, rows, columns, i, j int) {
	// fmt.Println(rows, columns, i, j)
	if i >= 0 && j >= 0 && i < rows && j < columns && grid[i][j] == '1' {
		grid[i][j] = '0'
		clear(grid, rows, columns, i-1, j)
		clear(grid, rows, columns, i, j-1)
		clear(grid, rows, columns, i+1, j)
		clear(grid, rows, columns, i, j+1)
	}
}

func main() {
	numIslands(
		[][]byte{
			[]byte{'1', '1', '1', '1', '0'},
			[]byte{'1', '1', '0', '1', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '0', '0', '0'},
		})

	numIslands(
		[][]byte{
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '1', '0', '0'},
			[]byte{'0', '0', '0', '1', '1'},
		})

	numIslands([][]byte{})
}
