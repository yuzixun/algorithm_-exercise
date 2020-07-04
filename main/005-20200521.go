package main

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	leftMax, rightMax := 0, 0
	// result := ""
	for i := 0; i < len(s); i++ {
		left1, right1 := strLength(s, i-1, i+1)
		left2, right2 := strLength(s, i, i+1)
		if right1-left1 < right2-left2 {
			left1, right1 = left2, right2
		}

		if rightMax-leftMax < right1-left1 {
			leftMax, rightMax = left1, right1
		}
	}

	// fmt.Println(s[leftMax : rightMax+1])
	return s[leftMax : rightMax+1]
}

func strLength(s string, left, right int) (int, int) {
	curLeft, curRight := 0, 0
	for ; left >= 0 && right < len(s); left, right = left-1, right+1 {
		if s[left] != s[right] {
			return curLeft, curRight
		}
		curLeft, curRight = left, right
	}
	return curLeft, curRight
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	longestPalindrome("babad")
	longestPalindrome("cbbd")
	longestPalindrome("1")
}
