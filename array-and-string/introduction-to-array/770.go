package main

import "fmt"

func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	cur := 0
	for i := 0; i < len(nums); i++ {
		if 2*cur+nums[i] == sum {
			return i
		}
		cur += nums[i]
	}

	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{}))
}
