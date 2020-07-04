package main

import (
	"strings"
)

func reverseWords(s string) string {
	arr := []string{}

	i, j := 0, 0
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			arr = append(arr, reverseWord(s, j, i))
			j = i + 1
		}
	}
	arr = append(arr, reverseWord(s, j, i))
	// fmt.Println(arr, len(arr))
	return strings.Join(arr, " ")
}

func reverseWord(s string, start, end int) string {
	// fmt.Printf("%+v\n", s)
	items := []byte(s[start:end])
	left, right := 0, len(items)-1
	for left < right {
		items[left], items[right] = items[right], items[left]
		left++
		right--
	}

	return string(items)
}

func main() {
	reverseWords("Let's take LeetCode contest")
}
