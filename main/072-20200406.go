package main

func minDistance(word1 string, word2 string) int {
	size1, size2 := len(word1), len(word2)
	dp := make([][]int, size1+1)
	for i := 0; i < size1+1; i++ {
		dp[i] = make([]int, size2+1)
	}

	for i := 0; i <= size1; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= size2; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= size1; i++ {
		for j := 1; j <= size2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	// fmt.Println(dp)
	return dp[size1][size2]
}

func min(a, b, c int) int {
	result := 0
	switch {
	case a <= b && a <= c:
		result = a
	case b <= a && b <= c:
		result = b
	case c <= a && c <= b:
		result = c
	}
	return result
}

func main() {
	minDistance("horse", "ros")
	minDistance("intention", "execution")
}
