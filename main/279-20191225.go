package main

import "fmt"

func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
