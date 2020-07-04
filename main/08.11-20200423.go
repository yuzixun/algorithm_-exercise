package main

func waysToChange(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	coins := []int{25, 10, 5, 1}
	for _, coin := range coins {
		for i := coin; i <= n; i++ {
			dp[i] += dp[i-coin]
			dp[i] %= 1000000007
		}
	}

	// fmt.Println(dp)
	return dp[n]
}

func main() {
	waysToChange(5)
	waysToChange(10)
	waysToChange(100)
}
