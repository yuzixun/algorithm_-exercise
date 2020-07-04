package main

func coinChange(coins []int, amount int) int {
	coinsSize := len(coins)
	if coinsSize == 0 {
		return -1
	}
	dp := make([]int, amount+1)
	NULL := 1 << 32
	for i := 1; i <= amount; i++ {
		dp[i] = NULL
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < coinsSize; j++ {
			if coins[j] <= i {
				// fmt.Println(coins[j], i, dp[i-coins[j]]+1, dp[i])
				dp[i] = min(dp[i-coins[j]]+1, dp[i])
			}
		}
		// fmt.Println(dp)
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
func main() {
	coinChange([]int{1, 2, 5}, 11)
	coinChange([]int{2}, 3)
}
