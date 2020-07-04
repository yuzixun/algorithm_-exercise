package main

import "fmt"

func containsDuplicate(nums []int) bool {
	cache := map[int]bool{}

	for _, num := range nums {
		if _, ok := cache[num]; ok {
			return true
		}
		cache[num] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
