package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			}
		}
	}

	// fmt.Println(result)
	return result
}

func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}
