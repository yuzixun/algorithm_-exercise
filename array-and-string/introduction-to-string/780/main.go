package main

import "fmt"

func strStr(haystack string, needle string) int {
	size1, size2 := len(haystack), len(needle)

	for i := 0; i < size1-size2+1; i++ {
		if haystack[i:i+size2] == needle {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("", ""))
}
