package main

func maximalRectangle(matrix [][]byte) int {
	maxArea := 0

	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			var width int

			if j > 0 {
				dp[i][j] = dp[i][j-1] + 1
				width = dp[i][j]
			} else {
				dp[i][j] = 1
				width = 1
			}

			for k := i; k >= 0; k-- {
				width = min(width, dp[k][j])
				maxArea = max(maxArea, width*(i-k+1))
			}
		}
	}

	// fmt.Println(maxArea)
	return maxArea
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	matrix := [][]byte{
// 		[]byte{'1', '0', '1', '0', '0'},
// 		[]byte{'1', '0', '1', '1', '1'},
// 		[]byte{'1', '1', '1', '1', '1'},
// 		[]byte{'1', '0', '0', '1', '0'},
// 	}

// 	maximalRectangle(matrix)
// }
