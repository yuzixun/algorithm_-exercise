package main

func firstMissingPositive(nums []int) int {
	foundZero := false
	for i := 0; i < len(nums); i++ {
		nums[i]--
		if nums[i] == 0 {
			foundZero = true
		}
	}

	if foundZero {
		return firstMissingPositive(nums) + 1
	}

	return 1
}

// func main() {
// 	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
// 	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
// 	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
// }
