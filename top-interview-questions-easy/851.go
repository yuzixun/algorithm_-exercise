package main

import "fmt"

func isPerfectSquare(num int) bool {
	x := num
	for x*x > num {
		x = (x + num/x) / 2
	}

	return x*x == num
}

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
	fmt.Println(isPerfectSquare(25))
	fmt.Println(isPerfectSquare(99))
	fmt.Println(isPerfectSquare(100))
	fmt.Println(isPerfectSquare(1000))
	fmt.Println(isPerfectSquare(64))
	fmt.Println(isPerfectSquare(1))
}
