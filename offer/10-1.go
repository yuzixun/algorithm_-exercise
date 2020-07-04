package main

func fib(n int) int {
	if n == 0 {
		return 0
	}
	a, b, result := 0, 1, 1
	for i := 1; i < n; i++ {
		result = (a + b) % 1000000007
		a, b = b, result
	}

	// fmt.Println(result)
	return result
}

// func main() {
// 	fib(0)
// 	fib(1)
// 	fib(2)
// 	fib(5)
// 	fib(93)
// }

// 755204270
// 965550172
// 720754435
// 686304600
// 407059028
// 93363621
// 500422649
// 593786270
// 94208912
// 687995182
// 965550172-720754435
