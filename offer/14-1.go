package main

import (
	"math"
)

func cuttingRope1(n int) int {
	// dp := []int{0, 1, 1}

	// for i := 3; i <= n; i++ {
	// 	dp[i%3] = max(max(max(dp[(i-1)%3], i-1), 2*max(dp[(i-2)%3], i-2)), 3*max(dp[(i-3)%3], i-3))
	// }

	// return dp[n%3]

	if n < 4 {
		return n - 1
	}
	if n == 4 {
		return 4
	}

	result := -1
	switch n % 3 {
	case 0:
		result = int(math.Pow(3, float64(n/3)))
	case 1:
		result = int(math.Pow(3, float64(n/3-1)) * 4)
	case 2:
		result = int(math.Pow(3, float64(n/3)) * 2)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
