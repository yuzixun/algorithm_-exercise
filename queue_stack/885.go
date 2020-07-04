package main

// func findTargetSumWays(nums []int, S int) int {
// 	count := 0
// 	dfs(nums, S, 0, 0, &count)
// 	// fmt.Println(count)
// 	return count
// }

// func dfs(nums []int, target, index, current int, count *int) {
// 	// fmt.Println(target, index, current, *count)
// 	if index == len(nums) {
// 		if target == current {
// 			(*count)++
// 		}

// 		return
// 	}

// 	dfs(nums, target, index+1, current-nums[index], count)
// 	dfs(nums, target, index+1, current+nums[index], count)
// }

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum < S || (sum+S)%2 != 0 {
		return 0
	}
	m := (sum + S) / 2
	dp := make([]int, len(nums))
	dp[0] = 1
	for _, num := range nums {
		for i := m; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}

	return dp[m]
}
