package main

import "fmt"

func firstMissingPositive(nums []int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		for nums[i] > 0 && nums[i] < size && nums[i] != nums[nums[i]-1] {
			index := nums[i] - 1
			nums[index], nums[i] = nums[i], nums[index]
		}
	}

	for i, num := range nums {
		if i+1 != num {
			return i + 1
		}
	}

	return size + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	fmt.Println(firstMissingPositive([]int{1, 1}))
}
