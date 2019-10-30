package main

import "fmt"

func maxProduct(nums []int) int {
	result, size := -1<<32, len(nums)
	maxData := make([]int, size, size)
	minData := make([]int, size, size)
	maxData[0] = nums[0]
	result = nums[0]
	for i := 1; i < size; i++ {
		minData[i] = min(min(maxData[i-1]*nums[i], minData[i-1]*nums[i]), nums[i])
		maxData[i] = max(max(maxData[i-1]*nums[i], minData[i-1]*nums[i]), nums[i])

		if result < maxData[i] {
			result = maxData[i]
		}
	}

	fmt.Println(maxData)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr := []int{-3}
	fmt.Println(maxProduct(arr))
}
