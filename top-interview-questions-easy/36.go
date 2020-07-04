package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	size := len(s)
	left, right := 0, size-1
	for left <= right {
		for left <= size-1 && !(s[left] >= 'a' && s[left] <= 'z' || s[left] >= '0' && s[left] <= '9') {
			left++
		}

		for right >= 0 && !(s[right] >= 'a' && s[right] <= 'z' || s[right] >= '0' && s[right] <= '9') {
			right--
		}

		if left > right {
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
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome(""))
	fmt.Println(isPalindrome("0P"))
	fmt.Println(isPalindrome(" "))
}
