package main

import "fmt"

func sumNums(n int) int {
	var helper func(sum *int, n int) bool
	helper = func(sum *int, n int) bool {
		*sum += n
		return n > 0 && helper(sum, n-1)
	}

	sum := 0
	helper(&sum, n)
	return sum
}

func main() {
	fmt.Println(sumNums(3))
	fmt.Println(sumNums(9))
}
