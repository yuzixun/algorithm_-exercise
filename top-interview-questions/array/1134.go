package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	// columns := len(matrix[0])

	row, column := 0, len(matrix[0])-1

	for row < rows && column >= 0 {
		// fmt.Println(row, column, matrix[row][column])
		if matrix[row][column] > target {
			column--
		} else if matrix[row][column] < target {
			row++
		} else {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{
		[]int{1, 4, 7, 11, 15},
		[]int{2, 5, 8, 12, 19},
		[]int{3, 6, 9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}, 5))

	fmt.Println(searchMatrix([][]int{
		[]int{1, 4, 7, 11, 15},
		[]int{2, 5, 8, 12, 19},
		[]int{3, 6, 9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}, 20))

}
