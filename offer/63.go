package main

import "fmt"

func maxProfit(prices []int) int {
	size := len(prices)
	dp := make([][][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = []int{[]int{0, 0}, []int{-prices[i], 0}}
	}

	for i := 1; i < size; i++ {

	}

	fmt.Println(dp)
	return dp[size-1]
	// minPrice, result := prices[0], -1
	// for i := 1; i < size; i++ {
	// 	minPrice = min(prices[i], minPrice)
	// 	result = max(result, prices[i]-minPrice)
	// }
	// fmt.Println(result)
	// return result
}

// func maxProfit(prices []int) int {
// 	if len(prices) == 0 {
// 		return 0
// 	}
// 	size := len(prices)
// 	minPrice, result := prices[0], 0
// 	for i := 1; i < size; i++ {
// 		minPrice = min(prices[i], minPrice)
// 		result = max(result, prices[i]-minPrice)
// 	}
// 	// fmt.Println(result)
// 	return result
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	maxProfit([]int{7, 1, 5, 3, 6, 4})
	maxProfit([]int{7, 6, 4, 3, 1})
}
