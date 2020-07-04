package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return []int{}
	}
	columns := len(matrix[0])
	size := rows * columns
	result := make([]int, size)
	left, right, top, end := 0, columns-1, 0, rows-1
	index := 0

	for {
		for i := left; i <= right; i++ {
			result[index] = matrix[top][i]
			index++
		}
		top++
		if index >= size {
			break
		}

		for i := top; i <= end; i++ {
			result[index] = matrix[i][right]
			index++
		}
		right--
		if index >= size {
			break
		}

		for i := right; i >= left; i-- {
			result[index] = matrix[end][i]
			index++
		}
		end--
		if index >= size {
			break
		}

		for i := end; i >= top; i-- {
			result[index] = matrix[i][left]
			index++
		}
		left++
		if index >= size {
			break
		}
	}

	// fmt.Println(result)
	return result
}

func main() {
	fmt.Println(spiralOrder([][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}))

	fmt.Println(spiralOrder([][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
	}))
}
