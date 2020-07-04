package main

func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	dp := make([]int, size)
	result := 1
	for i := 0; i < size; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
				result = max(dp[i], result)
			}
		}
	}
	// fmt.Println(dp)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
// 	lengthOfLIS(nums)
// }
