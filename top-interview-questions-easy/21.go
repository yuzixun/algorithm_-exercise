package main

func removeDuplicates(nums []int) int {
	size := len(nums)
	slow, fast := 0, 1
	for fast < size {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	// fmt.Println(nums, slow, fast)
	return slow + 1
}

// func main() {
// 	fmt.Println(removeDuplicates([]int{1}))
// 	fmt.Println(removeDuplicates([]int{1, 1}))
// 	fmt.Println(removeDuplicates([]int{0, 1, 2}))
// 	fmt.Println(removeDuplicates([]int{0, 0, 1, 2}))
// 	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
// }
