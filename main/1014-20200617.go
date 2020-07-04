package main

import "fmt"

func maxScoreSightseeingPair(A []int) int {
	pre := A[0] + 0
	result := 0

	for i := 1; i < len(A); i++ {
		result = max(result, pre+A[i]-i)
		pre = max(pre, A[i]+i)
	}

	// fmt.Println(result, pre)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
	fmt.Println(maxScoreSightseeingPair([]int{1, 5, 2, 6}))
	fmt.Println(maxScoreSightseeingPair([]int{1, 3, 5}))
}
