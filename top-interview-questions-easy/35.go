package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cache := make([]int, 26)

	for i := 0; i < len(s); i++ {
		cache[s[i]-'a']++
		cache[t[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if cache[i] != 0 {
			return false
		}
	}

	return true
}
func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
