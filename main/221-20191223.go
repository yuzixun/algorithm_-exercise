package main

func maximalSquare(matrix [][]byte) int {
	maxY := len(matrix)
	if maxY == 0 {
		return 0
	}
	maxX := len(matrix[0])

	dp := make([][]int, maxY)
	for i := 0; i < maxY; i++ {
		dp[i] = make([]int, maxX)
	}

	// for i := 0; i < maxX; i++ {
	// 	if matrix[0][i] == '1' {
	// 		dp[0][i] = 1
	// 	}
	// }

	// for i := 0; i < maxY; i++ {
	// 	if matrix[i][0] == '1' {
	// 		dp[i][0] = 1
	// 	}
	// }

	result := 0
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			if matrix[i][j] == '1' {
				if i > 0 && j > 0 && matrix[i][j-1] == '1' && matrix[i-1][j] == '1' {
					dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
				} else {
					dp[i][j] = 1
				}

				result = max(result, dp[i][j])
			}
		}
	}

	// fmt.Println(dp, result)
	return result * result
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

// func main() {
// 	matrix := [][]byte{
// 		[]byte{'0', '0', '0', '1'},
// 		[]byte{'1', '1', '0', '1'},
// 		[]byte{'1', '1', '1', '1'},
// 		[]byte{'0', '1', '1', '1'},
// 		[]byte{'0', '1', '1', '1'},
// 	}

// 	// matrix := [][]byte{
// 	// 	[]byte{'1'},
// 	// }
// 	maximalSquare(matrix)
// }
