package main

func rob(nums []int) int {
	preMax, curMax := 0, 0

	for i := 0; i < len(nums); i++ {
		tmp := curMax
		if preMax+nums[i] > curMax {
			curMax = preMax + nums[i]
		}
		preMax = tmp
	}
	return curMax
}

// func main() {
// 	// nums := []int{1, 2, 3, 1}
// 	// nums := []int{2, 7, 9, 3, 1}
// 	nums := []int{1, 2, 3}
// 	// nums := []int{}
// 	fmt.Println(rob(nums))
// }
