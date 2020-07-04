package main

func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	with, without := -prices[0], 0
	// dp := make([][]int, len(prices))

	// for i := 0; i < len(prices); i++ {
	// dp[i] = []int{0, -prices[i]}
	// }

	for i := 1; i < len(prices); i++ {
		with = max(with, without-prices[i])
		without = max(without, with+prices[i]-fee)
	}

	// fmt.Println(dp)
	// fmt.Println(without)
	return without
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	maxProfit([]int{1, 3, 2, 8, 4, 9}, 2)
// 	maxProfit([]int{1, 2}, 1)
// }
