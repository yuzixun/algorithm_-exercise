package main

func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// 天、交易次数、手上是否有股票
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, 3)
		for j := 0; j <= 2; j++ {
			dp[i][j] = []int{0, -prices[i]}
		}
	}

	for i := 1; i < len(prices); i++ {

		dp[i][0][0] = dp[i-1][0][0]
		dp[i][0][1] = max(dp[i-1][0][1], dp[i-1][0][0]-prices[i])

		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][0][1]+prices[i])
		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][1][0]-prices[i])

		dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][1][1]+prices[i])
	}

	// fmt.Println(dp)
	return dp[len(prices)-1][2][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4})
// 	maxProfit([]int{1, 2, 3, 4, 5})
// 	maxProfit([]int{7, 6, 4, 3, 1})
// }
