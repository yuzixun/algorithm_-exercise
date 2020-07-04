package main

func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}

	left, right := byte('('), byte(')')
	size := len(s)
	max := -1
	dp := make([]int, size)

	for i := 1; i < size; i++ {
		if s[i] == right {
			if s[i-1] == left {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}

			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == left {
				before := 0
				if i-dp[i-1]-2 >= 0 {
					before = dp[i-dp[i-1]-2]
				}
				dp[i] = dp[i-1] + 2 + before
			}

		}
	}

	for _, i := range dp {
		if max < i {
			max = i
		}
	}

	return max
}

// func main() {

// 	fmt.Println(longestValidParentheses("(()"))
// 	fmt.Println(longestValidParentheses(")()())"))
// }
