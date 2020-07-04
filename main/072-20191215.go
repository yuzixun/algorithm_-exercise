package main

func minDistance(word1 string, word2 string) int {
	size1, size2 := len(word1), len(word2)

	dp := make([][]int, size1+1)
	for i := 0; i <= size1; i++ {
		dp[i] = make([]int, size2+1)
	}

	for i := 1; i <= size1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	for j := 1; j <= size2; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}

	for i := 1; i <= size1; i++ {
		for j := 1; j <= size2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
			}
		}
	}

	return dp[size1][size2]
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	fmt.Println(minDistance("horse", "ros"))
// 	fmt.Println(minDistance("intention", "execution"))
// 	fmt.Println(minDistance("", ""))
// 	fmt.Println(minDistance("a", "a"))
// 	fmt.Println(minDistance("", "a"))
// }
