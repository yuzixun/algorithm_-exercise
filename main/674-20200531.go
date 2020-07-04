package main

func findLengthOfLCIS(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	result, cur := 0, 0

	last := -1 << 32
	for i := 0; i < size; i++ {
		// fmt.Println(i, nums[i], cur, result)
		if last < nums[i] {
			cur++
			result = max(result, cur)
		} else {
			cur = 1
		}

		last = nums[i]
	}

	// fmt.Println(result)
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func main() {
	findLengthOfLCIS([]int{1, 3, 5, 4, 7})
	findLengthOfLCIS([]int{2, 2, 2, 2, 2})
	findLengthOfLCIS([]int{1, 3, 5, 4, 2, 3, 4, 5})
}
