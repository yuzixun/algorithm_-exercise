package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {

	result := len(nums) + 1

	left, right, cur := 0, 0, 0
	for right < len(nums) {

		cur += nums[right]
		for cur >= s {
			if result > right-left+1 {
				result = right - left + 1
			}

			cur -= nums[left]
			left++
		}

		right++
	}

	if result == len(nums)+1 {
		return 0
	}

	return result
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(7, []int{9, 2, 3, 1, 2, 4, 3}))
}
