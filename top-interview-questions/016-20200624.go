package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	result, delta := 0, math.MaxFloat64

	for i := 0; i < len(nums); i++ {

		left, right := i+1, len(nums)-1
		for left < right {
			cur := nums[left] + nums[right] + nums[i]

			if cur > target {
				right--
			} else if cur < target {
				left++
			} else {
				return target
			}

			curDelta := math.Abs(float64(cur - target))
			if curDelta < float64(delta) {
				delta = curDelta
				result = cur
			}
		}
	}

	// fmt.Println(result)
	return result
}

func main() {
	threeSumClosest([]int{-1, 2, 1, -4}, 1)
}
