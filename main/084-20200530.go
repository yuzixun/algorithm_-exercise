package main

import "fmt"

func largestRectangleArea(heights []int) int {
	size := len(heights)

	left, right := make([]int, size), make([]int, size)
	for i := 0; i < size; i++ {
		right[i] = size
	}

	stack := []int{}
	for i := 0; i < size; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// fmt.Println(left, right)
	result := 0
	for i := 0; i < size; i++ {
		result = max(result, (right[i]-left[i]-1)*heights[i])
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
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
