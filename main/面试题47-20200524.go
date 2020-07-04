package main

func maxValue(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	dp := make([][]int, rows)
	columns := len(grid[0])
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < columns; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = max(dp[i][j], max(dp[i-1][j], dp[i][j-1])+grid[i][j])
		}
	}

	// fmt.Println(dp[rows-1][columns-1])
	return dp[rows-1][columns-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	maxValue([][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	})
}
