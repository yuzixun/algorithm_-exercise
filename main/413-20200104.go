package main

import "fmt"

func numberOfArithmeticSlices(A []int) int {
	size := len(A)

	current, sum := 0, 0
	for i := 2; i < size; i++ {

		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			current++
			sum += current
		} else {
			current = 0
		}
	}
	fmt.Println(sum)
	return sum
}

// func main() {
// 	numberOfArithmeticSlices([]int{1, 2, 3, 4})
// 	numberOfArithmeticSlices([]int{1, 3, 5, 7, 9})
// 	numberOfArithmeticSlices([]int{1, 1, 2, 5, 7})
// }
