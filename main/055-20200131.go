package main

func canJump(nums []int) bool {
	size := len(nums)
	k := 0
	// cache := make([]bool, size)

	// cache[size-1] = true
	// fmt.Println(cache)

	for i := 0; i < size; i++ {
		if i > k {
			return false
		}
		k = max(k, i+nums[i])
	}
	// fmt.Println(k)
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	canJump([]int{2, 3, 1, 1, 4})
// 	canJump([]int{3, 2, 1, 0, 4})
// 	canJump([]int{3, 2, 1, 10, 4})
// }
