package main

func findMaxForm(strs []string, m int, n int) int {
	if (m <= 0 && n <= 0) || len(strs) == 0 {
		return 0
	}

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	arrLen := len(strs)
	for a := 0; a < arrLen; a++ {
		zero, one := 0, 0
		str := strs[a]
		strLen := len(str)
		for b := 0; b < strLen; b++ {
			if str[b] == '0' {
				zero++
			} else {
				one++
			}
		}

		for i := m; i-zero >= 0; i-- {
			for j := n; j-one >= 0; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
			}
		}
	}

	// fmt.Println(dp)
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3)
// 	findMaxForm([]string{"10", "0", "1"}, 1, 1)
// }
