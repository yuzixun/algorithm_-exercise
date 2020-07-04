package main

import "fmt"

func maxProduct(nums []int) int {
	maxTmp, minTmp := nums[0], nums[0]

	maxResult := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			maxTmp = max(nums[i], maxTmp*nums[i])
			minTmp = min(nums[i], minTmp*nums[i])
		} else {
			maxTmp, minTmp = max(nums[i], minTmp*nums[i]), min(nums[i], maxTmp*nums[i])
		}

		maxResult = max(maxResult, maxTmp)
		// fmt.Println(maxResult, maxTmp, minTmp)
	}

	return maxResult
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{-4, -3, -2}))
}
