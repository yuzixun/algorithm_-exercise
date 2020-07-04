package main

import "fmt"

func strStr(haystack string, needle string) int {
	size1, size2 := len(haystack), len(needle)
	if len(needle) == 0 {
		return 0
	}
	if size1 < size2 {
		return -1
	}
	for i := 0; i < size1; i++ {
		j := 0
		for ; j < size2 && i+j < size1; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == size2 {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
}
