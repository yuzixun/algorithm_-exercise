package main

import "fmt"

func rob(nums []int) int {
	preMax, curMax := 0, 0
	for _, num := range nums {
		tmp := curMax
		if preMax+num > curMax {
			curMax = preMax + num
		}
		preMax = tmp
	}
	if preMax > curMax {
		return preMax
	}
	return curMax
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
