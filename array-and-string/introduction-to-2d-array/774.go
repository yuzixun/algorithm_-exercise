package main

func findDiagonalOrder(matrix [][]int) []int {
	result := []int{}
	rows := len(matrix)
	if rows == 0 {
		return result
	}

	columns := len(matrix[0])
	cur := 0
	i, j := 0, 0
	up := true
	for i+j == cur && cur <= rows+columns-2 {
		// fmt.Println(result, i, j)
		result = append(result, matrix[i][j])

		if (i == 0 || j == columns-1) && up {
			if j == columns-1 {
				i++
			} else {
				j++
			}
			cur++
			up = !up
		} else if (j == 0 || i == rows-1) && !up {
			if i == rows-1 {
				j++
			} else {
				i++
			}
			cur++
			up = !up
		} else {
			if up {
				i--
				j++
			} else {
				i++
				j--
			}
		}

	}
	// fmt.Println(result)
	return result
}

func main() {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	findDiagonalOrder(matrix)
}
