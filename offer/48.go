package main

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	slow, fast, result := 0, 0, 1

	for ; fast < len(s); fast++ {
		for i := fast - 1; i >= slow; i-- {
			// fmt.Println("-------------", s[i], s[fast], slow, fast)
			if s[i] == s[fast] {
				slow = i + 1
				break
			}
		}
		result = max(result, fast-slow+1)
		// fmt.Println(slow, fast, result)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
// 	fmt.Println(lengthOfLongestSubstring("bbbbb"))
// 	fmt.Println(lengthOfLongestSubstring("pwwkew"))
// 	fmt.Println(lengthOfLongestSubstring(""))
// 	fmt.Println(lengthOfLongestSubstring("au"))
// }
