package main

func dailyTemperatures(T []int) []int {
	stack := []int{}
	result := make([]int, len(T))
	for i := 0; i < len(T); i++ {

		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			day := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[day] = i - day
		}
		stack = append(stack, i)
	}
	// fmt.Println(result)
	return result
}

// func main() {
// 	dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
// }
