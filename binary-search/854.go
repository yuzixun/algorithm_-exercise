package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left < right {
		mid = left + (right-left)/2

		if nums[mid] == nums[right] {
			right--
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{2, 2, 2, 0, 1, 1, 1}))
	fmt.Println(findMin([]int{2, 2, 2, 1, 1, 1}))
	fmt.Println(findMin([]int{1, 3, 5}))
	fmt.Println(findMin([]int{1, 3, 3}))
	fmt.Println(findMin([]int{3, 1, 3, 3}))
	fmt.Println(findMin([]int{3, 3, 3, 3, 1, 3, 3}))
}
