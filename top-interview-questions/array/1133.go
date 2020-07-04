package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	small, mid := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
		if num <= small {
			small = num
		} else if num <= mid {
			mid = num
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println(increasingTriplet([]int{3, 4, 1, 2, 5}))
	fmt.Println(increasingTriplet([]int{3, 3, 1, 2, 5}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
}
