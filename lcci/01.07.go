package main

func rotate(matrix [][]int) {
	size := len(matrix)

	if size <= 1 {
		return
	}

	for i := 0; i < size; i++ {
		a := size - 1 - i
		for j := i; j < a; j++ {
			b := size - 1 - j
			matrix[i][j], matrix[j][a], matrix[a][b], matrix[b][i] = matrix[b][i], matrix[i][j], matrix[j][a], matrix[a][b]
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
