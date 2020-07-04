package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	for n%3 == 0 {
		n /= 3
		if n == 1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(0))
	fmt.Println(isPowerOfThree(9))
	fmt.Println(isPowerOfThree(45))
}
