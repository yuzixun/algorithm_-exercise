package main

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := -1 << 32
	tmpMin, tmpMax := result*-1, result
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			tmpMax, tmpMin = max(tmpMin*nums[i], nums[i]), min(tmpMax*nums[i], nums[i])
		} else {
			tmpMax, tmpMin = max(tmpMax*nums[i], nums[i]), min(tmpMin*nums[i], nums[i])
		}

		result = max(result, tmpMax)
	}

	// fmt.Println(result, tmpMax, tmpMin)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	maxProduct([]int{2, 3, -2, 4})
	maxProduct([]int{-2, 0, -1})
	maxProduct([]int{-2})
}
