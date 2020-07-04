package main

import "fmt"

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]

		fast = nums[fast]
		fast = nums[fast]
	}

	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 1}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
}
