package main

func new21Game(N int, K int, W int) float64 {
	dp := make([]float64, K+W)

	r := 0.0

	for i := K; i < K+W; i++ {
		if i <= N {
			dp[i] = 1
		} else {
			dp[i] = 0
		}

		r += dp[i]
	}

	for i := K - 1; i >= 0; i-- {
		dp[i] = r / float64(W)
		r = r - dp[i+W] + dp[i]
	}

	// fmt.Println(dp)
	return dp[0]
}

func main() {
	new21Game(10, 1, 10)
	new21Game(6, 1, 10)
	new21Game(21, 17, 10)

}
