package main

import "fmt"

func isPalindrome(s string) bool {
	if s == " " {
		return true
	}

	var str []byte
	for i := 0; i < len(s); i++ {
		if 'a' <= s[i] && s[i] <= 'z' || '0' <= s[i] && s[i] <= '9' {
			str = append(str, s[i])
		}

		if 'A' <= s[i] && s[i] <= 'Z' {
			str = append(str, s[i]-'A'+'a')
		}
	}

	left, right := 0, len(str)-1
	for left < right {
		if str[left] == str[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}
