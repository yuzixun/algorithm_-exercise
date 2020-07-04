package main

import "fmt"

func findTheLongestSubstring(s string) int {
	result := 0
	size := len(s)
	hash := map[byte]int{'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}

	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = []int{0, 0}
	}

	for i := 0; i < size; i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			hash[s[i]]++
			add := true
			for _, c := range hash {
				if c%2 != 0 {
					add = false
					break
				}
			}
			if add {
				dp[i][1] = max(dp[i-1][0]+1, dp[i-1][1])
				// dp[i][0] = dp[i-1][0]
			} else {
				if i == 0 {
					dp[i][0] = 1
				} else {
					dp[i][0] = max(dp[i-1][1]+1, dp[i-1][0])
					// dp[i][1] = dp[i-1][1]
				}
			}
		} else {
			if i == 0 {
				dp[i][1] = 1
				dp[i][0] = 1
			} else {
				dp[i][1] = dp[i-1][1] + 1
				dp[i][0] = dp[i-1][0] + 1
			}
		}
	}

	fmt.Println(dp)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	findTheLongestSubstring("eleetminicoworoep")
	findTheLongestSubstring("leetcodeisgreat")
	findTheLongestSubstring("bcbcbc")
}
