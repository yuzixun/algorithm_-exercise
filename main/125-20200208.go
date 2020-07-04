package main

func isPalindrome(s string) bool {
	if s == " " {
		return true
	}

	var str []byte
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			str = append(str, s[i])
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			str = append(str, s[i]-'A'+'a')
		}
	}

	left, right := 0, len(str)-1
	for left <= right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// func main() {
// 	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
// 	fmt.Println(isPalindrome("race a car"))
// 	// fmt.Println(byte("race a car"[0]))
// 	// fmt.Println(byte("A "[1]))
// }
