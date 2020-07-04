package main

import (
	"math"
)

func printNumbers(n int) []int {
	max := math.Pow10(n)
	arr := make([]int, int(max)-1)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	// fmt.Println(arr)
	return arr
}

// func main() {
// 	printNumbers(1)
// 	printNumbers(2)
// 	printNumbers(3)
// }
