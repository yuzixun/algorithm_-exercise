package main

func handleKInf(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		profit += max(prices[i]-prices[i-1], 0)
	}
	return profit
}

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	if k > len(prices)/2 {
		return handleKInf(prices)
	}

	dp := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = []int{0, -prices[0]}
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j < k+1; j++ {
			dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i])
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])
		}
	}

	// fmt.Println(dp)

	result := 0
	for i := 0; i < k+1; i++ {
		result = max(result, dp[i][0])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(maxProfit(2, []int{2, 4, 1}))
// 	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
// 	fmt.Println(maxProfit(2, []int{2, 1, 4, 5, 2, 9, 7}))
// 	fmt.Println(maxProfit(2, []int{3, 3, 5, 0, 0, 3, 1, 4}))
// 	fmt.Println(maxProfit(1, []int{1, 2}))
// }
