package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = []int{0, -prices[i]}
	}

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		if i == 1 {
			dp[i][1] = max(dp[i-1][1], -prices[i])
		} else {
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
		}
	}
	// fmt.Println(dp)
	result := 0
	for i := 0; i < len(prices); i++ {
		result = max(result, dp[i][0])
	}
	// fmt.Println(result)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	maxProfit([]int{1, 2, 3, 0, 2})
// 	maxProfit([]int{1, 2})
// 	maxProfit([]int{2})
// }
