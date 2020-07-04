package main

import "fmt"

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	fmt.Println(left)
	return left
}

func main() {
	findPeakElement([]int{1, 2, 3, 1})
	findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})
}
