package main

func massage(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	a, b := 0, nums[0]
	for i := 1; i < size; i++ {
		ta := max(a, b)
		tb := a + nums[i]

		a, b = ta, tb
	}

	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(massage([]int{1, 2, 3}))
// 	fmt.Println(massage([]int{1, 2, 3, 1}))
// 	fmt.Println(massage([]int{2, 7, 9, 3, 1}))
// 	fmt.Println(massage([]int{2, 1, 4, 5, 3, 1, 1, 3}))
// }
