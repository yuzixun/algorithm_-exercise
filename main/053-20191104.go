package main

import "fmt"

func maxSubArray(nums []int) int {
	size := len(nums)
	result := make([]int, size, size)
	result[0] = nums[0]

	maxValue := nums[0]
	for i := 1; i < size; i++ {
		result[i] = max(result[i-1]+nums[i], nums[i])
		maxValue = max(maxValue, result[i])
	}

	fmt.Println(result)
	return maxValue
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
// 	fmt.Println(maxSubArray(arr))
// }
