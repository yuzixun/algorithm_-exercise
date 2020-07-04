package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// cur := nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if result < nums[i] {
			result = nums[i]
		}
	}
	// fmt.Println(result)
	return result
}

// func main() {
// 	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
// 	maxSubArray([]int{-2, -1})
// 	maxSubArray([]int{-2})
// }
