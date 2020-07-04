package main

import "fmt"

func canThreePartsEqualSum(A []int) bool {

	sum := 0
	for _, item := range A {
		sum += item
	}

	if sum%3 != 0 {
		return false
	}

	target := sum / 3

	sum, count := 0, 0
	for _, item := range A {
		sum += item
		if sum == target {
			count++
			sum = 0
		}
	}

	return count >= 3
}

func main() {
	fmt.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}))
	fmt.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}))
	fmt.Println(canThreePartsEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))
	fmt.Println(canThreePartsEqualSum([]int{6, 1, 1, 13, -1, 0, -10, 20}))

}
