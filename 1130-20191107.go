package main

import "fmt"

func mctFromLeafValues(arr []int) int {
	size := len(arr)
	dp := make([][]int, size, size)
	maxArr := make([][]int, size, size)

	for i := range dp {
		maxArr[i] = make([]int, size, size)
		dp[i] = make([]int, size, size)
	}

	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			maxVal := 0
			for k := i; k <= j; k++ {
				maxVal = max(maxVal, arr[k])
			}
			fmt.Println(i, j, maxVal)

			maxArr[i][j] = maxVal

			if i == j {
				dp[i][j] = 0
			} else {
				dp[i][j] = 1 << 32
			}
		}
	}

	fmt.Println(dp, maxArr)
	for i := 1; i < size; i++ {
		for j := 0; j < size-i; j++ {
			e := i + j
			for k := j; k < e; k++ {
				dp[j][e] = min(dp[j][e], dp[j][k]+dp[k+1][e]+maxArr[j][k]*maxArr[k+1][e])
			}
		}
	}
	fmt.Println(dp)
	return dp[0][size-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(mctFromLeafValues([]int{6, 2, 4}))
}
