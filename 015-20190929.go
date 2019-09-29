package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	size := len(nums)
	sort.Ints(nums)
	for i := 0; i < size; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		num := nums[i]
		left, right := i+1, size-1
		for left < right {
			current := num + nums[left] + nums[right]
			if current == 0 {
				result = append(result, []int{num, nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if current > 0 {
				right--
			} else if current < 0 {
				left++
			}
		}
	}
	return result
}

func contains(array *[]int, element int) bool {
	for _, number := range *array {
		if number == element {
			return true
		}
	}
	return false
}

func main() {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{-1, -1, 0, 1, 2}
	// nums := []int{0, 0, 0, 0}
	// nums := []int{-2, 0, 1, 1, 2}
	nums := []int{-1, -1, 0, 1, 2}

	fmt.Println(nums)
	fmt.Println(threeSum(nums))
}
