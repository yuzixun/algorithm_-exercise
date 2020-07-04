package main

import (
	"strings"
)

func reverseStr(s string, k int) string {
	size := len(s)

	for i := 0; i < size; {
		var endPos int
		if i+k > size {
			endPos = size
		} else {
			endPos = i + k
		}
		s = strings.Replace(s, s[i:endPos], reverseString(s[i:endPos]), 1)
		i += 2 * k
	}
	return s
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// func main() {
// 	s := "abcdefgabcdefg"
// 	k := 2
// 	fmt.Println(reverseStr(s, k))
// }
