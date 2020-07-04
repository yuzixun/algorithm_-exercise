package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	size := len(s)
	if size == 0 {
		return true
	}

	left, right := 0, size-1
	for left < right {
		for left < right && !('a' <= s[left] && s[left] <= 'z' || '0' <= s[left] && s[left] <= '9') {
			left++
		}

		for left < right && !('a' <= s[right] && s[right] <= 'z' || '0' <= s[right] && s[right] <= '9') {
			right--
		}

		if left >= right {
			return true
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}
