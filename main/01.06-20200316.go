package main

import (
	"fmt"
	"strconv"
)

func compressString(S string) string {
	arr := []byte{}
	prev := S[0]
	count := 1
	for i := 1; i < len(S); i++ {
		if S[i] != prev {
			arr = append(arr, prev)
			arr = append(arr, strconv.Itoa(count)...)
			prev = S[i]
			count = 1
		} else {
			count++
		}
	}
	arr = append(arr, prev)
	arr = append(arr, strconv.Itoa(count)...)
	if len(S) < len(arr) {
		return S
	}
	return string(arr)
}

func main() {
	fmt.Println(compressString("aabcccccaaa"))
	fmt.Println(compressString("abbccd"))
}
