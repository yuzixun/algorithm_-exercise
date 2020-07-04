package main

import "fmt"

// func trap(height []int) int {
// 	size := len(height)
// 	if size == 0 {
// 		return 0
// 	}
// 	result := 0
// 	left := make([]int, size)
// 	right := make([]int, size)

// 	left[0] = height[0]
// 	right[size-1] = height[size-1]

// 	for i := 1; i < size; i++ {
// 		left[i] = max(left[i-1], height[i-1])
// 	}

// 	for i := size - 2; i >= 0; i-- {
// 		right[i] = max(right[i+1], height[i+1])
// 	}

// 	// fmt.Println(left, right)
// 	for i := 0; i < size; i++ {
// 		cur := min(left[i], right[i]) - height[i]
// 		if cur > 0 {
// 			result += cur
// 		}

// 	}

// 	return result
// }

func trap(height []int) int {
	size := len(height)
	maxLeft, maxRight, result := 0, 0, 0
	left, right := 1, size-2

	for i := 1; i < size-1; i++ {
		if height[left-1] < height[right+1] {
			maxLeft = max(maxLeft, height[left-1])
			cur := maxLeft - height[left]
			if cur > 0 {
				result += cur
			}
			left++
		} else {
			maxRight = max(maxRight, height[right+1])
			cur := maxRight - height[right]
			if cur > 0 {
				result += cur
			}
			right--
		}
	}

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
func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{}))
}
