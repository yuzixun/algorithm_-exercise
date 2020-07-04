package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	cur, reversed := x, 0
	for cur != 0 {
		reversed = reversed*10 + cur%10
		cur = cur / 10
	}

	return reversed == x
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
