package main

import "fmt"

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	result := -1 << 32
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(result, dp[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, 1, 2, 1, -5, 4}))
}
