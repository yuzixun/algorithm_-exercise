package main

import "fmt"

func strStr(haystack string, needle string) int {
	hSize := len(haystack)
	nSize := len(needle)

	for i := 0; i <= hSize-nSize; i++ {
		j := 0
		for ; j < nSize; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == nSize {
			return i
		}
	}
	return -1
}

func main() {
	haystack := ""
	needle := ""

	fmt.Println("strStr is ", strStr(haystack, needle))
}
