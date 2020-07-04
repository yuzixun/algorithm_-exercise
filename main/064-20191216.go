package main

import "fmt"

func minPathSum(grid [][]int) int {
	ii := len(grid)

	jj := len(grid[0])
	if ii == 0 || jj == 0 {
		return 0
	}
	dp := make([][]int, ii)
	for i := 0; i < ii; i++ {
		dp[i] = make([]int, jj)
	}

	for i := 0; i < ii; i++ {
		for j := 0; j < jj; j++ {
			switch true {
			case i == 0 && j == 0:
				dp[i][j] = grid[i][j]
			case i == 0:
				dp[i][j] = dp[i][j-1] + grid[i][j]
			case j == 0:
				dp[i][j] = dp[i-1][j] + grid[i][j]
			default:
				dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
			}
		}
	}
	fmt.Println(dp[ii-1][jj-1])
	return dp[ii-1][jj-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// func main() {
// 	grid := [][]int{{}}
// 	// grid := [][]int{{1, 3, 3}, {1, 5, 1}, {4, 2, 1}}
// 	minPathSum(grid)
// }
