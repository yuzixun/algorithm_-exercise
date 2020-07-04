package main

import (
	"strings"
)

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	// fmt.Println("s is", s)
	i, j := len(s)-1, len(s)-1

	arr := []string{}
	for i >= 0 {
		for i >= 0 && s[i] != ' ' {
			i--
		}
		arr = append(arr, s[i+1:j+1])
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
	}

	// fmt.Println(arr)
	return strings.Join(arr, " ")
}

func main() {
	reverseWords("the sky is blue")
	reverseWords("  hello world!  ")
	reverseWords("a good   example")
}
