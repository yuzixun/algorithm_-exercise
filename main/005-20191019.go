package main

// func longestPalindrome(s string) string {
// 	result := ""
// 	size := len(s)
// 	for i := 0; i < size; i++ {
// 		curMax := stringCompare(s, i, i+1)
// 		s2 := stringCompare(s, i, i)

// 		if len(curMax) < len(s2) {
// 			curMax = s2
// 		}

// 		if len(result) < len(curMax) {
// 			result = curMax
// 		}

// 	}
// 	return result
// }

// func stringCompare(s string, left, right int) string {
// 	size := len(s)

// 	result := ""
// 	for ; left >= 0 && right <= size-1 && s[left] == s[right]; left, right = left-1, right+1 {
// 		result = s[left : right+1]
// 	}
// 	return result
// }

func longestPalindrome(s string) string {
	resultLeft, resultRight := 0, 0
	size := len(s)
	for i := 0; i < size; i++ {
		p := i
		left1, right1 := stringCompare(s, i-1, p)
		left2, right2 := stringCompare(s, i, i)

		if right1-left1 < right2-left2 {
			left1, right1 = left2, right2
		}

		if resultRight-resultLeft < right1-left1 {
			resultLeft, resultRight = left1, right1
		}

	}
	return s[resultLeft:resultRight]
}

func stringCompare(s string, left, right int) (int, int) {
	size := len(s)
	resultLeft, resultRight := 0, 0

	for ; left >= 0 && right <= size-1 && s[left] == s[right]; left, right = left-1, right+1 {
		resultLeft, resultRight = left, right+1
	}
	return resultLeft, resultRight
}

// func main() {
// 	// s := "a"
// 	s := "cbbd"
// 	fmt.Println(longestPalindrome(s))
// }
