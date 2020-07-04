package main

import (
	"strings"
)

func helper(s string, k int, start, end int) int {
	if start > end {
		return 0
	}
	size := len(s)
	hash := make(map[byte]int, size)

	// fmt.Println(start, end, s, s[start:end])
	for i := start; i < end; i++ {
		hash[s[i]]++
	}

	// fmt.Println(hash)
	for key, value := range hash {
		if value > 0 && value < k {
			pos := strings.IndexByte(s[start:end], key) + start
			// fmt.Println(start, value, end, s[start:end], pos)
			return max(helper(s, k, start, pos), helper(s, k, pos+1, end))
		}
	}

	return end - start
}

func longestSubstring(s string, k int) int {
	// result := helper(s, k, 0, len(s))
	return helper(s, k, 0, len(s))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	// fmt.Println(longestSubstring("aaabbb", 3))
// 	// fmt.Println(longestSubstring("aaabb", 3))
// 	// fmt.Println(longestSubstring("ababbc", 2))
// 	// fmt.Println(longestSubstring("ababacb", 3))
// 	// fmt.Println(longestSubstring("weitong", 2))
// }
