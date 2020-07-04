package main

import "fmt"

func numSquares(n int) int {
	if n <= 2 {
		return n
	}

	MAX := 1 << 32
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		cur := MAX
		for j := 1; j*j <= i; j++ {
			cur = min(cur, dp[i-j*j])
		}
		dp[i] = cur + 1
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
	fmt.Println(numSquares(14))
	fmt.Println(numSquares(15))
	fmt.Println(numSquares(16))
}
