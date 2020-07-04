package main

import "fmt"

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return fast
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println(findDuplicate([]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}))
	fmt.Println(findDuplicate([]int{75, 75, 75, 75, 75, 91, 75, 75, 75, 75, 75, 75, 30, 75, 75, 78, 75, 75, 75, 75, 75, 7, 79, 93, 75, 75, 15, 75, 75, 75, 75, 75, 75, 4, 75, 75, 21, 75, 75, 19, 75, 59, 75, 76, 75, 75, 75, 75, 75, 75, 75, 33, 75, 75, 75, 58, 75, 75, 5, 75, 97, 81, 48, 42, 75, 87, 75, 75, 25, 27, 94, 20, 75, 75, 75, 29, 75, 75, 86, 67, 75, 75, 75, 65, 75, 75, 75, 75, 75, 39, 75, 56, 75, 75, 75, 75, 3, 75, 75, 75}))
}
