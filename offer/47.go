package main

func maxValue(grid [][]int) int {
	rowSize := len(grid)
	colSize := len(grid[0])
	if rowSize == 0 {
		return 0
	}
	dp := make([][]int, rowSize)
	for i := 0; i < rowSize; i++ {
		dp[i] = make([]int, colSize)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < rowSize; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < colSize; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < rowSize; i++ {
		for j := 1; j < colSize; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	// fmt.Println(dp)
	return dp[rowSize-1][colSize-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	maxValue([][]int{
// 		[]int{1, 3, 1},
// 		[]int{1, 5, 1},
// 		[]int{4, 2, 1}})

// 	maxValue([][]int{
// 		[]int{1, 2, 5},
// 		[]int{3, 2, 1}})
// }
