package main

import "fmt"

func jump(nums []int) int {
	last := 0
	maxPos := 0
	result := 0

	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, i+nums[i])
		if last == i {
			last = maxPos
			result++
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
