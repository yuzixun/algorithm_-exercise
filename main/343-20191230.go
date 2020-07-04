package main

import (
	"math"
)

// n % 3  == 0 直接返回 3 **(n/3)
// n % 3  == 1 直接返回 3 ** ((n-1)/3) * 4
// n % 3  == 2 直接返回 3 **(n/3) * 2
func integerBreak(n int) int {
	switch true {
	case n <= 3:
		return n - 1
	case n%3 == 0:
		return int(math.Pow(3, float64(n/3)))
	case n%3 == 1:
		return int(math.Pow(3, float64(n/3)-1) * 4)
	case n%3 == 2:
		return int(math.Pow(3, float64(n/3)) * 2)
	}
	return 1
}

// func main() {
// 	fmt.Println(integerBreak(58))
// 	fmt.Println(integerBreak(10))
// 	fmt.Println(integerBreak(2))
// 	fmt.Println(integerBreak(3))
// }
