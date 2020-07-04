package main

func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	rows := len(matrix)
	// fmt.Println(rows, columns)
	for i := 0; i < rows/2; i++ {
		// fmt.Println(i, rows-i)
		matrix[i], matrix[rows-1-i] = matrix[rows-1-i], matrix[i]
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// fmt.Println(matrix)
}

// func main() {
// 	rotate([][]int{
// 		[]int{1, 2, 3},
// 		[]int{4, 5, 6},
// 		[]int{7, 8, 9},
// 	})

// 	rotate([][]int{
// 		[]int{5, 1, 9, 11},
// 		[]int{2, 4, 8, 10},
// 		[]int{13, 3, 6, 7},
// 		[]int{15, 14, 12, 16},
// 	})
// }
