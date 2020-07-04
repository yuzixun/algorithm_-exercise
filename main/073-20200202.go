package main

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	rows, columns := len(matrix), len(matrix[0])
	row0, column0 := false, false

	for i := 0; i < columns; i++ {
		if matrix[0][i] == 0 {
			row0 = true
			break
		}
	}

	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			column0 = true
			break
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if row0 {
		for i := 0; i < columns; i++ {
			matrix[0][i] = 0
		}
	}

	if column0 {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
	// fmt.Println(matrix)
}

// func main() {
// 	setZeroes([][]int{
// 		[]int{1, 1, 1},
// 		[]int{1, 0, 1},
// 		[]int{1, 1, 1},
// 	})

// 	setZeroes([][]int{
// 		[]int{0, 1, 2, 0},
// 		[]int{3, 4, 5, 2},
// 		[]int{1, 3, 1, 5},
// 	})
// }
