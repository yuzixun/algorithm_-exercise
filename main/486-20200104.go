package main

import "fmt"

func PredictTheWinner(nums []int) bool {
	size := len(nums)

	if size&1 == 0 {
		return true
	}
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
	}

	for l := 1; l < size; l++ {
		for i := 0; i < size-l; i++ {
			j := i + l
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}

	fmt.Println(dp)
	return dp[0][size-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	PredictTheWinner([]int{1, 5, 2})
// 	PredictTheWinner([]int{1, 5, 233, 7})
// 	PredictTheWinner([]int{})
// }
