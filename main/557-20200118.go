package main

import (
	"fmt"
)

func reverseWords(s string) string {

	slow, fast := 0, 0
	bArr := []byte(s)
	for slow <= len(bArr)-1 {
		// fmt.Println(fast, bArr[fast])

		for len(bArr)-1 >= fast && bArr[fast] != 32 {
			fast++
		}

		left, right := slow, fast-1
		for left < right {
			bArr[left], bArr[right] = bArr[right], bArr[left]
			left++
			right--
		}

		fast++
		slow = fast
	}

	fmt.Println(string(bArr))
	return string(bArr)
}
