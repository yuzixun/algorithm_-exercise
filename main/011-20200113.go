package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := -1

	for left <= right {
		result = max(result, min(height[left], height[right])*(right-left))

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/* func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
*/
