package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	y := 0
	for x > y {
		y = y*10 + x%10
		x /= 10
		// fmt.Println(x, y)
	}

	return x == y || x == y/10
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
