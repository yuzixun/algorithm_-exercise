package main

import (
	"fmt"
)

func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 0; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
}
