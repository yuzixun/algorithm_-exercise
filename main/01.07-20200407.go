package main

func rotate(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}

	columns := len(matrix[0])
	for i := 0; i < rows/2; i++ {
		for j := i; j < columns-1-i; j++ {
			matrix[i][j], matrix[j][columns-1-i], matrix[columns-1-i][rows-1-j], matrix[rows-1-j][i] = matrix[rows-1-j][i], matrix[i][j], matrix[j][columns-1-i], matrix[columns-1-i][rows-1-j]
		}
	}

}

func main() {
	rotate([][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9}})

	rotate([][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16}})
}
