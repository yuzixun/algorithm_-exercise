package main

func longestValidParentheses(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}

	dp := make([]int, size)
	result := 0
	for i := 1; i < size; i++ {
		if s[i] == '(' {
			continue
		}

		if s[i-1] == '(' {
			dp[i] = 2

			if i >= 2 {
				dp[i] += dp[i-2]
			}
		} else if s[i-1] == ')' {
			ii := i - dp[i-1] - 1
			if ii >= 0 && s[ii] == '(' {
				dp[i] = dp[i-1] + 2

				if ii >= 1 {
					dp[i] += +dp[ii-1]
				}
			}
		}

		result = max(result, dp[i])
	}
	// fmt.Println(result, dp)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	longestValidParentheses("(()")
	longestValidParentheses(")()())")
	longestValidParentheses("()(())")
	longestValidParentheses("(())(())")
}
