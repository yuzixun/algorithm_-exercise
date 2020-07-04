package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	cur := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		cur = max(cur+nums[i], nums[i])
		result = max(cur, result)
	}

	// fmt.Println(result)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}
