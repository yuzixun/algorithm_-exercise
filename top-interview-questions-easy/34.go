package main

import "fmt"

func firstUniqChar(s string) int {
	cache := make([]int, 26)

	for i := 0; i < len(s); i++ {
		cache[s[i]-'a']++
	}

	for i, c := range s {
		if cache[c-97] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}
