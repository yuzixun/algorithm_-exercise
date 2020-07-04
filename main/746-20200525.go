package main

func minCostClimbingStairs(cost []int) int {
	size := len(cost)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return cost[0]
	}

	// cur1, cur2 := cost[0], cost[1]
	dp := make([]int, size)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < size; i++ {
		dp[i] = min(dp[i-1]+cost[i], dp[i-2]+cost[i])
	}

	// fmt.Println(dp)
	return min(dp[size-1], dp[size-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minCostClimbingStairs([]int{10, 15, 20})
	minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
}
