package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	// result := []int{}
	count, maxCount := 0, len(matrix)*len(matrix[0])
	result := make([]int, maxCount)

	for {
		for i := left; i <= right; i++ {
			result[count] = matrix[up][i]
			count++
		}
		if maxCount <= count {
			break
		}
		up++

		for i := up; i <= down; i++ {
			result[count] = matrix[i][right]
			count++
		}
		if maxCount <= count {
			break
		}
		right--

		for i := right; i >= left; i-- {
			result[count] = matrix[down][i]
			count++
		}
		if maxCount <= count {
			break
		}
		down--

		for i := down; i >= up; i-- {
			result[count] = matrix[i][left]
			count++
		}
		if maxCount <= count {
			break
		}
		left++
	}
	// fmt.Println(result)
	return result
}

// func main() {
// 	spiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}})
// 	spiralOrder([][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}})
// }
