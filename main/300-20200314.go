package main

import "fmt"

func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	result := []int{nums[0]}
	left, right, mid := 0, 0, 0
	for i := 1; i < size; i++ {
		left, right = 0, len(result)-1
		if result[right] < nums[i] {
			result = append(result, nums[i])
		} else {
			for left <= right {
				mid = left + (right-left)/2
				if result[mid] < nums[i] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			result[left] = nums[i]
			// fmt.Println(result, left, right, mid)
		}
	}

	return len(result)
}

func main() {

	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 8, 101, 18}))
	fmt.Println("-------------------")
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 4}))
	fmt.Println(lengthOfLIS([]int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}))
}
