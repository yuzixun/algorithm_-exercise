package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	counter := map[int]int{}
	result := []int{}

	for _, num := range nums1 {
		counter[num]++
	}

	for _, num := range nums2 {
		if counter[num] > 0 {
			counter[num]--
			result = append(result, num)
		}
	}

	return result
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
