package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	size := len(s)
	result := make([]bool, size+1)

	result[0] = true
	for i := 0; i <= size; i++ {
		for j := 0; j < i; j++ {
			if !result[j] {
				continue
			}

			for _, word := range wordDict {
				if word == s[j:i] {
					result[i] = true
					break
				}
			}
		}
	}
	return result[size]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
