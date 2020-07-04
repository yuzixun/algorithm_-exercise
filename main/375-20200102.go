package main

// dp[i][j] = x + max(dp[i][x-1], dp[x+1][j])
func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	for i := n; i >= 1; i-- {
		for j := i; j <= n; j++ {
			if i == j {
				dp[i][j] = 0
			} else {
				dp[i][j] = 1 << 32
				for x := i; x <= j; x++ {
					dp[i][j] = min(dp[i][j], max(dp[i][x-1], dp[x+1][j])+x)
				}
			}
		}
	}
	// fmt.Println(dp)
	return dp[1][n]
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

// func main() {
// 	getMoneyAmount(10)
// }
