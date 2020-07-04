package main

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1, amount+1)
	NULL := 1 << 32
	for i := 0; i <= amount; i++ {
		dp[i] = NULL
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}
	if dp[amount] == NULL {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func main() {
// 	// fmt.Println(coinChange([]int{1, 2, 5}, 11))
// 	fmt.Println(coinChange([]int{1, 2, 5}, 20))
// 	// fmt.Println(coinChange([]int{2}, 3))

// }
