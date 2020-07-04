package main

func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	// fmt.Println(dp)
	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	maxProfit([]int{7, 1, 5, 3, 6, 4})
// 	maxProfit([]int{7, 6, 4, 3, 1})
// }

// 如果还没有购买，则
// dp[i][j] = max(dp[i])
