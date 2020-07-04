package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	r := myPow(x, n/2)
	if n%2 == 0 {
		return r * r
	} else {
		return r * r * x
	}
}

// func main() {
// 	fmt.Println(myPow(2, 1))
// }
