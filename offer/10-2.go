package main

func numWays(n int) int {
	if n == 1 {
		return 1
	}
	a, b, result := 1, 2, 2

	for i := 3; i <= n; i++ {
		result = (a + b) % 1000000007
		a, b = b, result
	}

	return result
}

// func main() {
// 	fmt.Println(numWays(1))
// 	fmt.Println(numWays(2))
// 	fmt.Println(numWays(7))
// }
