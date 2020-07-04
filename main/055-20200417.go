package main

import "fmt"

func canJump(nums []int) bool {
	size := len(nums)
	k := 0
	for i := 0; i < size; i++ {
		if i > k {
			return false
		}
		k = max(k, nums[i]+i)
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
