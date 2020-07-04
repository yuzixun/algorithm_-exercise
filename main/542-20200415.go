package main

func updateMatrix(matrix [][]int) [][]int {

	rows := len(matrix)
	if rows == 0 {
		return [][]int{}
	}
	NULL := 1 << 32
	result := make([][]int, rows)

	columns := len(matrix[0])
	for i := 0; i < rows; i++ {
		result[i] = make([]int, columns)
		for j := 0; j < columns; j++ {
			result[i][j] = NULL
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == 0 {
				result[i][j] = 0
				continue
			}
			if i > 0 {
				result[i][j] = min(result[i][j], result[i-1][j]+1)

			}
			if j > 0 {
				result[i][j] = min(result[i][j], result[i][j-1]+1)

			}
		}
	}

	for i := rows - 1; i >= 0; i-- {
		for j := columns - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				result[i][j] = 0
				continue
			}
			if i < rows-1 {
				result[i][j] = min(result[i][j], result[i+1][j]+1)
			}
			if j < columns-1 {
				result[i][j] = min(result[i][j], result[i][j+1]+1)
			}
		}
	}
	// fmt.Println(result)
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	updateMatrix(
		[][]int{[]int{0, 0, 0},
			[]int{0, 1, 0},
			[]int{0, 0, 0}})

	updateMatrix(
		[][]int{[]int{0, 0, 0},
			[]int{0, 1, 0},
			[]int{1, 1, 1}})

}
