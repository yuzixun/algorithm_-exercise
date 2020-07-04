package main

func maxProduct(nums []int) int {
	size := len(nums)

	result := -1 << 32
	curMin, curMax := 1, 1

	for i := 0; i < size; i++ {
		if nums[i] < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(curMax*nums[i], nums[i])
		curMin = max(curMin*nums[i], nums[i])

		result = max(curMax, result)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func maxProduct(nums []int) int {
// 	// result := -1 << 32
// 	size := len(nums)

// 	dpMax := make([]int, size)
// 	dpMin := make([]int, size)
// 	dpMax[0] = nums[0]
// 	dpMin[0] = nums[0]
// 	for i := 1; i < size; i++ {
// 		dpMax[i] = max(max(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]), nums[i])
// 		dpMin[i] = min(min(dpMin[i-1]*nums[i], dpMax[i-1]*nums[i]), nums[i])
// 	}

// 	result := -1 << 32
// 	for _, v := range dpMax {
// 		result = max(result, v)
// 	}
// 	// fmt.Println(dpMax, result)
// 	return result
// }

func main() {
	maxProduct([]int{2, 3, -2, 4})
	maxProduct([]int{-2, 0, -1})
}
