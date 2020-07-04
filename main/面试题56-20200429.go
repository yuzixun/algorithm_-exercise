package main

import "fmt"

func singleNumbers(nums []int) []int {
	tmp := 0
	for _, num := range nums {
		tmp ^= num
	}

	div := 1
	for i := 0; i < 32; i++ {
		if div&tmp != 0 {
			break
		}
		div <<= 1
	}

	num1, num2 := 0, 0
	for _, num := range nums {
		if num&div != 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}

	return []int{num1, num2}
}

func main() {
	fmt.Println(singleNumbers([]int{4, 1, 4, 6}))
	fmt.Println(singleNumbers([]int{1, 2, 10, 4, 1, 4, 3, 3}))
	fmt.Println(singleNumbers([]int{}))
}
