package main

func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	symbol := 1
	if n < 0 {
		symbol = -1
		n *= -1
	}

	res := n % 2

	c := myPow(x, n/2)

	var result float64

	result = float64(c * c)

	if res == 1 {
		result *= x
	}

	if symbol == -1 {
		result = 1 / result
	}

	return result
}

// func main() {
// 	fmt.Println(myPow(2.00000, 10))
// 	fmt.Println(myPow(2.10000, 3))
// 	fmt.Println(myPow(2.00000, -2))
// }
