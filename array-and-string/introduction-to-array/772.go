package main

import "fmt"

func plusOne(digits []int) []int {
	carry := 1
	var cur int
	for i := len(digits) - 1; i >= 0; i-- {
		cur = digits[i] + carry
		digits[i] = cur % 10
		carry = cur / 10
		if carry == 0 {
			break
		}
	}

	if carry == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9, 9}))
}
