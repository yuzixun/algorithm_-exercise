package main

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

// func main() {
// 	// array := []int{1, 3, 4, 2, 2}
// 	array := []int{3, 1, 3, 4, 2}
// 	fmt.Println("findDuplicate ", findDuplicate(array))
// }
