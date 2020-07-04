package main

import (
	"strings"
)

// func reverseWords(s string) string {
// 	arr := strings.Fields(s)
// 	slow, fast := 0, len(arr)-1
// 	for slow < fast {
// 		arr[slow], arr[fast] = arr[fast], arr[slow]
// 		slow++
// 		fast--
// 	}
// 	return strings.Join(arr, " ")
// }

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
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
	return strings.Join(arr, " ")
}

func main() {
	reverseWords("the sky is blue")
	reverseWords("  hello world!  ")
	reverseWords("a good   example")
}
