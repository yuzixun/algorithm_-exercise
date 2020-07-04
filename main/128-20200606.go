package main

import "fmt"

func longestConsecutive(nums []int) int {
	helper := map[int]int{}

	cur, result := 0, 0

	for _, num := range nums {
		_, ok := helper[num]
		if ok {
			continue
		}

		left, right := helper[num-1], helper[num+1]

		cur = left + 1 + right
		result = max(result, cur)

		helper[num-left] = cur
		helper[num] = cur
		helper[num+right] = cur
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
