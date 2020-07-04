package main

import "fmt"

func maxSubArray(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	result := nums[0]
	last := nums[0]
	for i := 1; i < size; i++ {
		last = max(nums[i], last+nums[i])
		result = max(result, last)
	}
	fmt.Println(result)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	maxSubArray([]int{-2, 1, -3, 4, -1, -1, 2, 1, -5, 4})
	maxSubArray([]int{})
	maxSubArray([]int{-1})
}
