package main

import "fmt"

func dominantIndex(nums []int) int {
	big1, big2 := -1, -1
	bigIndex := -1
	for index, num := range nums {
		if num > big2 {
			big2 = num
		}

		if num > big1 {
			big1, big2 = num, big1
			bigIndex = index
		}
	}

	if big1 >= big2*2 {
		return bigIndex
	}
	return -1
}

func main() {
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}))
	fmt.Println(dominantIndex([]int{1, 2, 3, 4}))
	fmt.Println(dominantIndex([]int{0, 0, 3, 2}))
	fmt.Println(dominantIndex([]int{1}))
}
