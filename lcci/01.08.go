package main

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	firstX, firstY := false, false
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != 0 {
				continue
			}

			if i == 0 {
				firstX = true
			}
			if j == 0 {
				firstY = true
			}

			matrix[i][0] = 0
			matrix[0][j] = 0
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstX {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	if firstY {
		for i := 0; i < len(matrix); i++ {
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
