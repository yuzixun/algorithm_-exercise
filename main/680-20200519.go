package main

import "fmt"

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] == s[right] {
			left++
			right--
			continue
		}

		return isPalindrome(s[left+1:right+1]) || isPalindrome(s[left:right])
	}

	return true
}

func isPalindrome(s string) bool {
	// fmt.Println(s)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("abccbba"))
	fmt.Println(validPalindrome(""))
	fmt.Println(validPalindrome("abdda"))
}
