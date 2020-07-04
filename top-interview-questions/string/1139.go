package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	cache := map[string][]string{}

	return dfs(s, wordDict, &cache)
}

func dfs(s string, wordDict []string, cache *map[string][]string) []string {
	strs, ok := (*cache)[s]
	if ok {
		return strs
	}

	result := []string{}
	if len(s) == 0 {
		result = append(result, "")
		return result
	}

	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			list := dfs(s[len(word):], wordDict, cache)

			for _, item := range list {
				if len(item) == 0 {
					result = append(result, word)
				} else {
					result = append(result, word+" "+item)
				}
			}
		}
	}

	(*cache)[s] = result
	return result
}

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	fmt.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
