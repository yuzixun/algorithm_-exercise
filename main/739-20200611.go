package main

func dailyTemperatures(T []int) []int {
	size := len(T)
	if size == 0 {
		return []int{}
	}

	result := make([]int, size)
	stack := []int{}
	for index, item := range T {
		for len(stack) > 0 && T[stack[len(stack)-1]] < item {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[last] = index - last
		}

		stack = append(stack, index)
	}
	// fmt.Println(result)
	return result
}

func main() {
	dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
}
