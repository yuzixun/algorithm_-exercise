package main

import "fmt"

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}

	var mid int
	left, right := 0, len(nums)
	for left < right {
		mid = left + (right-left)/2

		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}

	if left >= len(nums) || nums[left] != target {
		return result
	}
	result[0] = left
	left, right = 0, len(nums)
	for left < right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	result[1] = left - 1
	return result
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{5, 5}, 6))
	fmt.Println(searchRange([]int{5}, 5))
}
