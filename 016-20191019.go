package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var result int
	var minDelta = math.MaxUint32

	size := len(nums)
	for i := 0; i < size; i++ {
		left, right := i+1, size-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			var delta int
			if sum > target {
				delta = sum - target
			} else {
				delta = target - sum
			}

			// fmt.Println(nums[i], nums[left], nums[right], delta, minDelta)
			if minDelta > int(delta) {
				result = sum
				minDelta = delta
			}

			if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	return result
}

// func main() {
// 	// target, nums := 1, []int{-1, 2, 1, -4}
// 	target, nums := -100, []int{1, 1, 1, 0}

// 	fmt.Println(threeSumClosest(nums, target))
// }
