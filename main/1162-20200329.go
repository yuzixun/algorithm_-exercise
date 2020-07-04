package main

import "fmt"

func maxDistance(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	columns := len(grid[0])
	dp := make([][]int, rows)
	NULL := 1 << 32
	result := -1
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
		for j := 0; j < columns; j++ {
			dp[i][j] = NULL
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if i > 0 {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j])
			}

			if j > 0 {
				dp[i][j] = min(dp[i][j-1]+1, dp[i][j])
			}

			result = min(result, dp[i][j])
		}
	}

	for i := rows - 1; i >= 0; i-- {
		for j := columns - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}

			if i < rows-1 {
				dp[i][j] = min(dp[i+1][j]+1, dp[i][j])
			}

			if j < columns-1 {
				dp[i][j] = min(dp[i][j+1]+1, dp[i][j])
			}

			result = max(result, dp[i][j])
		}
	}
	// fmt.Println(dp)
	if result == NULL {
		return -1
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(maxDistance([][]int{[]int{1, 0, 1}, []int{0, 0, 0}, []int{1, 0, 1}}))
	fmt.Println(maxDistance([][]int{[]int{1, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}))
}
