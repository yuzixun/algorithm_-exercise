package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, column := 0, len(matrix[0])-1
	if target < matrix[0][0] || target > matrix[len(matrix)-1][column] {
		return false
	}
	for row < len(matrix) && column >= 0 {
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
