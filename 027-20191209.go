package main

func removeElement(nums []int, val int) int {
	size := len(nums)
	result := 0
	for ii := 0; ii < size; ii++ {
		if nums[ii] != val {
			nums[result] = nums[ii]
			result++
		}
	}
	return result
}

// func main() {
// 	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
// 	val := 2
// 	fmt.Println(removeElement(nums, val))
// }
