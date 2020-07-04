package main

func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}

	res := 1

	for n > 4 {
		n -= 3
		res = (res * 3) % 1000000007
	}

	return (n * res) % 1000000007
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(cuttingRope(2))
// 	fmt.Println(cuttingRope(10))
// 	fmt.Println(cuttingRope(30))
// 	fmt.Println(cuttingRope(1000))
// }

// // 620946522
