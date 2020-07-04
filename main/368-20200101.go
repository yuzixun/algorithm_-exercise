package main

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	size := len(nums)
	if size == 0 {
		return []int{}
	}
	sort.Ints(nums)

	dp := make([]int, size)
	parent := make([]int, size)
	m, mi := 0, 0
	for i := 0; i < size; i++ {
		for j := i; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && dp[i] <= dp[j] {
				dp[i] = dp[j] + 1
				parent[i] = j

				if dp[i] > m {
					m = dp[i]
					mi = i
				}
			}
		}
	}

	result := make([]int, m)
	for i := 0; i < m; i++ {
		result[i] = nums[mi]
		mi = parent[mi]
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

// func main() {
// 	largestDivisibleSubset([]int{1, 2, 3})
// 	largestDivisibleSubset([]int{1, 2, 4, 8})
// }
