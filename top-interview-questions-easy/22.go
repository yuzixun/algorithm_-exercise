package main

func maxProfit(prices []int) int {
	// result := 0
	size := len(prices)
	if size == 0 {
		return 0
	}

	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = []int{0, -prices[i]}
	}

	for i := 1; i < size; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[size-1][0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
// 	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
// 	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
// }
