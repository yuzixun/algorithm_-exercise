package main

import (
	"fmt"
	"sort"
)

func isStraight(nums []int) bool {
	sort.Ints(nums)
	prev := 0
	zeroCount := 0
	for i := 0; i < 5; i++ {
		if nums[i] == 0 {
			zeroCount++
		}
		if prev == 0 {
			prev = nums[i]
			continue
		}
		if prev == nums[i] {
			return false
		}
		needZero := nums[i] - prev - 1
		zeroCount -= needZero
		prev = nums[i]
	}

	fmt.Println(zeroCount)

	return zeroCount >= 0
}

// func main() {
// 	isStraight([]int{1, 2, 3, 4, 5})
// 	isStraight([]int{0, 0, 1, 2, 5})
// 	isStraight([]int{0, 0, 1, 2, 3})
// 	isStraight([]int{0, 0, 1, 2, 7})
// 	isStraight([]int{0, 13, 10, 11, 12})
// 	isStraight([]int{0, 0, 2, 2, 5})
// }
