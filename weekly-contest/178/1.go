package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	size := len(nums)
	result := make([]int, size)

	// var count int
	for i := 0; i < size; i++ {
		count := 0
		for j := 0; j < size; j++ {
			// fmt.Println(nums[i], nums[j])
			if nums[i] > nums[j] {
				count++
			}
		}
		// fmt.Println("count is ", i, count)
		result[i] = count
	}

	fmt.Println(result)
	return result
}

// func main() {
// 	smallerNumbersThanCurrent([]int{6, 5, 4, 8})
// 	smallerNumbersThanCurrent([]int{7, 7, 7, 7})
// 	smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3})
// 	smallerNumbersThanCurrent([]int{})

// }
