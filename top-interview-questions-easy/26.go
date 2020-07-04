package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	hash := map[int]int{}

	for _, num := range nums1 {
		hash[num]++
	}

	result := []int{}

	for i := 0; i < len(nums2); i++ {
		v, ok := hash[nums2[i]]
		if ok && v > 0 {
			hash[nums2[i]]--
			result = append(result, nums2[i])
		}
	}

	fmt.Println(result)
	return result
}

func main() {
	intersect([]int{1, 2, 2, 1}, []int{2, 2})
	intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
}
