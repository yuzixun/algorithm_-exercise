package main

func rob(nums []int) int {
	return max(rob1(nums[1:]), rob1(nums[:len(nums)-1]))
}

func rob1(nums []int) int {
	// fmt.Println(nums)
	preMax, curMax := 0, 0
	for i := 0; i < len(nums); i++ {
		temp := curMax
		curMax = max(preMax+nums[i], curMax)
		preMax = temp
	}

	return curMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(rob([]int{2, 3, 2, 1, 5}))
// 	fmt.Println(rob([]int{1, 2, 3, 1}))
// 	fmt.Println(rob([]int{1, 2}))
// }
