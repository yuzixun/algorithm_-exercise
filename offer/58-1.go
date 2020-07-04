package main

import (
	"strings"
)

func reverseWords(s string) string {
	// if s == "" {
	// 	return ""
	// }

	// size := 0
	// result := ""
	// for i := len(s) - 1; i >= 0; i-- {
	// 	if s[i] == ' ' && size != 0 {
	// 		result = result + s[i+1:i+size+1] + " "
	// 		size = 0
	// 		continue
	// 	}

	// 	if s[i] != ' ' {
	// 		size++
	// 	}
	// }

	// if size > 0 {
	// 	result = result + s[0:size] + " "
	// }

	// if len(result) > 0 {
	// 	result = result[0 : len(result)-1]
	// }
	strs := strings.Split(s, " ")
	result := ""
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] == "" {
			continue
		}
		result = result + strs[i] + " "
	}
	return result
}
