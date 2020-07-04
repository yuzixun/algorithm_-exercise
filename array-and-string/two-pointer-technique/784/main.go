package main

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	result := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		result += nums[i]
	}
	// fmt.Println(result)
	return result
}

func main() {
	arrayPairSum([]int{1, 4, 3, 2})
}
