package main

func removeDuplicates(nums []int) int {
	size := len(nums)
	i, j := 0, 0
	for ; i < size && j < size; j++ {
		if nums[i] < nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	// fmt.Println(nums, i+1)
	return i + 1
}

// func main() {
// 	removeDuplicates([]int{1, 1, 2})
// 	removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
// }
