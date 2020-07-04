package main

import "fmt"

func rob(nums []int) int {
	pre, cur := 0, 0

	for _, num := range nums {
		tmp := cur
		if pre+num > cur {
			cur = pre + num
		}
		pre = tmp
	}

	return cur
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 1, 1, 2}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

// func rob(nums []int) int {
// 	size := len(nums)
// 	if size == 0 {
// 		return 0
// 	}
// 	if size == 1 {
// 		return nums[0]
// 	}
// 	if size == 2 {
// 		return max(nums[0], nums[1])
// 	}

// 	dp := make([]int, size)
// 	dp[0] = nums[0]
// 	dp[1] = nums[1]

// 	for i := 2; i < size; i++ {
// 		t1, t2 := -1, -1
// 		if i-2 >= 0 {
// 			t1 = dp[i-2] + nums[i]
// 		}
// 		if i-3 >= 0 {
// 			t2 = dp[i-3] + nums[i]
// 		}
// 		dp[i] = max(t1, t2)
// 	}

// 	return max(dp[size-1], dp[size-2])
// }
