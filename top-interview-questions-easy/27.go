package main

import "fmt"

func plusOne(digits []int) []int {
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		carry = digits[i] / 10
		digits[i] %= 10
		if carry == 0 {
			break
		}
	}
	if carry == 1 {
		return append([]int{carry}, digits...)
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{4, 3, 2, 9}))
	fmt.Println(plusOne([]int{9}))
}
