package main

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left]&1 == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left++
		}
	}
	// fmt.Println(nums)
	return nums
}

// func main() {
// 	exchange([]int{1, 2, 3, 4})
// 	exchange([]int{1, 2, 3, 5, 7})
// 	exchange([]int{11, 9, 3, 7, 16, 4, 2, 0})
// }
