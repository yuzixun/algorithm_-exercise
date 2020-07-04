package main

import "fmt"

func myPow(x float64, n int) float64 {
	var result float64
	if n == 0 {
		return 1
	}

	positive := true
	if n < 0 {
		positive = false
		n *= -1
	}
	odd := n%2 == 1
	cur := myPow(x, n/2)
	result = cur * cur

	if odd {
		result = x * result
	}
	if !positive {
		result = (1.0 / result)
	}
	return result
}

func main() {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, -2))
}
