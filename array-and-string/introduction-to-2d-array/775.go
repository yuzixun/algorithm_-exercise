package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	rows := len(matrix)
	if rows == 0 {
		return result
	}
	columns := len(matrix[0])
	curLen, length := 0, rows*columns
	cMin, cMax, rMin, rMax := 0, columns-1, 0, rows-1
	i, j := 0, 0
	for {
		for i, j = rMin, cMin; j <= cMax; j++ {
			result = append(result, matrix[i][j])
			curLen++
		}
		rMin++
		if curLen == length {
			break
		}

		for i, j = rMin, cMax; i <= rMax; i++ {
			result = append(result, matrix[i][j])
			curLen++
		}
		cMax--
		if curLen == length {
			break
		}

		for i, j = rMax, cMax; j >= cMin; j-- {
			result = append(result, matrix[i][j])
			curLen++
		}
		rMax--
		if curLen == length {
			break
		}

		for i, j = rMax, cMin; i >= rMin; i-- {
			result = append(result, matrix[i][j])
			curLen++
		}
		cMin++
		if curLen == length {
			break
		}
	}
	return result
}

func main() {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(spiralOrder(matrix))
	matrix = [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(matrix))
}
