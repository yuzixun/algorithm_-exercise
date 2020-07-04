package main

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	NULL := 1 << 32
	for i := 1; i <= amount; i++ {
		dp[i] = NULL
	}
	for i := 0; i < len(coins); i++ {
		dp[coins[i]] = 1
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] == NULL {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	coinChange([]int{1, 2, 5}, 11)
	coinChange([]int{2}, 3)
}
