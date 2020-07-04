package main

func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}

	columns := len(matrix[0])

	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
	}

	result := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			if i > 0 && j > 0 {
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
			} else {
				dp[i][j] = 1
			}

			result = max(result, dp[i][j])
		}
	}

	// fmt.Println(result)
	return result * result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	maximalSquare(
		[][]byte{
			[]byte{'1', '0', '1', '0', '0'},
			[]byte{'1', '0', '1', '1', '1'},
			[]byte{'1', '1', '1', '1', '1'},
			[]byte{'1', '0', '0', '1', '0'}})
}
