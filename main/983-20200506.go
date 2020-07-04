package main

func mincostTickets(days []int, costs []int) int {
	size := days[len(days)-1] + 1
	dp := make([]int, size)
	dayIndex := 0
	for i := 1; i < size; i++ {
		if days[dayIndex] == i {
			dp[i] = min(min(dp[max(0, i-1)]+costs[0], dp[max(0, i-7)]+costs[1]), dp[max(0, i-30)]+costs[2])
			dayIndex++
		} else {
			dp[i] = dp[i-1]
		}

	}

	// fmt.Println(dp[size-1])
	return dp[size-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15})
	mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15})
}
