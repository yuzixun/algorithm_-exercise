package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var results [][]int
	sort.Ints(nums)
	calc(target, 4, nums, []int{}, &results)
	return results
}

func calc(target, kCount int, nums, result []int, results *[][]int) {
	size := len(nums)
	for i := 0; i < size-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < size-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, size-1
			for left < right {
				r := nums[i] + nums[j] + nums[left] + nums[right]
				switch {
				case r == target:
					*results = append(*results, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					for left < right && nums[left] == nums[left-1] {
						left++
					}

				case r < target:
					left++
				case r > target:
					right--
				}
			}
		}
	}
}

func main() {
	// nums := []int{-1, 0, -5, -2, -2, -4, 0, 1, -2}
	// target := -9
	nums := []int{0, 0, 0, 0}
	target := 0
	fmt.Println(fourSum(nums, target))
}

// func calc(target, kCount int, nums, result []int, results *[][]int) {
// 	if len(nums) < kCount || kCount < 2 {
// 		return
// 	}

// 	if kCount == 2 {
// 		for i, j := 0, len(nums)-1; i < j; {

// 			if target == nums[i]+nums[j] {
// 				tmp := append(result, nums[i], nums[j])
// 				*results = append(*results, tmp)
// 				i++
// 				for i < j && nums[i] == nums[i-1] {
// 					i++
// 				}
// 			}

// 			if target > nums[i]+nums[j] {
// 				i++
// 			}

// 			if target < nums[i]+nums[j] {
// 				j--
// 			}
// 		}
// 	} else {
// 		for i, num := range nums {
// 			var items []int
// 			copy(items, nums[:i])
// 			items = append(items, nums[i+1:]...)
// 			if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
// 				calc(target-num, kCount-1, items, append(result, num), results)
// 			}
// 		}
// 	}
// }
