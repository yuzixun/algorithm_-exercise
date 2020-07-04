package main

import (
	"math"
	"strconv"
)

func countDigitOne(n int) int {
	if n == 0 {
		return 0
	}

	return helperCount(strconv.Itoa(n))
}

func helperCount(n string) int {
	if n == "0" {
		return 0
	}

	size := len(n)
	if size == 1 {
		return 1
	}

	firstCount := 0

	if n[0] == '1' {
		n, _ := strconv.Atoi(n[1:])
		firstCount = n + 1
	} else {
		firstCount = int(math.Pow10(size - 1))
	}

	midCount := int(n[0]-'0') * (size - 1) * int(math.Pow10(size-2))

	return firstCount + midCount + helperCount(n[1:])
}

// func main() {
// 	fmt.Println(countDigitOne(12))
// 	fmt.Println(countDigitOne(13))
// }
