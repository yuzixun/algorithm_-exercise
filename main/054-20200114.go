package main

func spiralOrder(matrix [][]int) []int {
	maxX := len(matrix)
	if maxX == 0 {
		return []int{}
	}

	maxY := len(matrix[0])
	result := make([]int, maxY*maxX)
	index := 0
	left, right, up, down := 0, maxY-1, 0, maxX-1

	for {
		// fmt.Println("left, right, up, down", left, right, up, down)
		for i := left; i <= right; i++ {
			result[index] = matrix[up][i]
			index++
		}
		up++
		if up > down {
			break
		}

		// fmt.Println("left, right, up, down", left, right, up, down)
		for i := up; i <= down; i++ {
			result[index] = matrix[i][right]
			index++
		}
		right--
		if right < left {
			break
		}

		// fmt.Println("left, right, up, down", left, right, up, down)
		for i := right; i >= left; i-- {
			result[index] = matrix[down][i]
			index++
		}
		down--
		if up > down {
			break
		}

		// fmt.Println("left, right, up, down", left, right, up, down)
		for i := down; i >= up; i-- {
			result[index] = matrix[i][left]
			index++
		}
		left++
		if right < left {
			break
		}
	}
	// fmt.Println(result)

	return result
}

// func main() {
// 	spiralOrder([][]int{
// 		[]int{1, 2, 3},
// 		[]int{4, 5, 6},
// 		[]int{7, 8, 9},
// 	})

// 	spiralOrder([][]int{
// 		[]int{1, 2, 3, 4},
// 		[]int{5, 6, 7, 8},
// 		[]int{9, 10, 11, 12},
// 	})
// }
