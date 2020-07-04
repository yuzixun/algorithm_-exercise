package main

import (
	"sort"
)

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, nums[len(nums)-1]-nums[0]

	var mid, count int
	for left < right {
		mid = left + (right-left)/2
		count = 0
		for i := 0; i < len(nums); i++ {
			j := 0
			for nums[i]-nums[j] > mid {
				j++
			}
			count += (i - j)
		}

		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// fmt.Println("left is", left)
	return left
}

func main() {
	smallestDistancePair([]int{1, 3, 1}, 1)
	smallestDistancePair([]int{1, 3, 1}, 2)
	smallestDistancePair([]int{1, 3, 1, 3, 3, 3}, 2)
	smallestDistancePair([]int{62, 100, 4}, 2)
}
