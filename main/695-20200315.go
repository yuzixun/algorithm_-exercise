package main

func maxAreaOfIsland(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	columns := len(grid[0])

	cur, max := 0, 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			cur = helper(grid, i, j, rows, columns)
			if cur > max {
				max = cur
			}
		}
	}

	return max
}

func helper(grid [][]int, i, j, rows, columns int) int {
	if i < 0 || i >= rows || j < 0 || j >= columns || grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0
	result := 1 + helper(grid, i-1, j, rows, columns) + helper(grid, i+1, j, rows, columns) + helper(grid, i, j-1, rows, columns) + helper(grid, i, j+1, rows, columns)
	return result
}

func main() {
	maxAreaOfIsland([][]int{{0, 0, 0, 0, 0, 0, 0, 0}})
	maxAreaOfIsland([][]int{
		[]int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	})

	maxAreaOfIsland([][]int{
		[]int{1, 1, 0, 0, 0},
		[]int{1, 1, 0, 0, 0},
		[]int{0, 0, 0, 1, 1},
		[]int{0, 0, 0, 1, 1},
	})

}
