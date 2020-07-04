package main

import "fmt"

// func twoSum(numbers []int, target int) []int {
// 	size := len(numbers)
// 	for i := 0; i < len(numbers); i++ {
// 		findTarget := target - numbers[i]

// 		left, right := i+1, size-1
// 		for left <= right {
// 			mid := left + (right-left)/2
// 			if numbers[mid] == findTarget {
// 				return []int{i + 1, mid + 1}
// 			} else if numbers[mid] < findTarget {
// 				left = mid + 1
// 			} else if numbers[mid] > findTarget {
// 				right = mid - 1
// 			}
// 		}
// 	}
// 	return []int{-1, -1}
// }

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		c := numbers[left] + numbers[right]
		if c == target {
			return []int{left + 1, right + 1}
		} else if c > target {
			right--
		} else if c < target {
			left++
		}
	}
	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 26))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 27))
	fmt.Println(twoSum([]int{1, 2, 3, 4, 4, 9, 56, 90}, 8))

}
