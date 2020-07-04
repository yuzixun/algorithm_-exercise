package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}

	if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	res := 1.0
	if n%2 == 1 {
		res *= x
		n--
	}

	c := myPow(x, n/2)
	return c * c * res
}

func main() {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, -2))
	fmt.Println(myPow(-2.00000, -2))
	fmt.Println(myPow(-2.00000, 3))

}
