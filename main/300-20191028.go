package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	result := -1 << 32
	dp := make([]int, size)
	for i := 0; i < size; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}
	return result
}

// func main() {
// 	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
// 	fmt.Println(lengthOfLIS(arr))
// }
