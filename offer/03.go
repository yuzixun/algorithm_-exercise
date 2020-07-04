package main

func findRepeatNumber(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}

			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	return 0
}

// func main() {
// 	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 5, 3}))
// }
