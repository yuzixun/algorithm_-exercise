package main

func maxSumDivThree(nums []int) int {
	dp := make([]int, 3)

	for i := 0; i < len(nums); i++ {
		mod := nums[i] % 3

		a := dp[(3+0-mod)%3]
		b := dp[(3+1-mod)%3]
		c := dp[(3+2-mod)%3]

		if a != 0 || mod == 0 {
			dp[0] = max(dp[0], a+nums[i])
		}
		if b != 0 || mod == 1 {
			dp[1] = max(dp[1], b+nums[i])
		}
		if c != 0 || mod == 2 {
			dp[2] = max(dp[2], c+nums[i])
		}
	}
	// fmt.Println(dp)
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	maxSumDivThree([]int{3, 6, 5, 1, 8})
// 	maxSumDivThree([]int{4})
// 	maxSumDivThree([]int{1, 2, 3, 4, 4})
// }
