package main

import (
	"math"
)

func reverse(x int) int {
	result := 0

	for x != 0 {
		mod := x % 10

		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && mod > math.MaxInt32%10) {
			return 0
		}

		if result < math.MinInt32/10 || (result == math.MinInt32/10 && mod < math.MinInt32%10) {
			return 0
		}

		result = result*10 + mod
		x = x / 10
	}

	return result
}

// func main() {
// 	fmt.Println(reverse(123))
// 	fmt.Println(reverse(-123))
// 	fmt.Println(reverse(120))
// 	fmt.Println(reverse(1534236469))
// 	fmt.Println(reverse(1563847412))
// }
