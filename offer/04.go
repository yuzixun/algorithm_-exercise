package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	x, y := 0, len(matrix[0])-1
	for x < len(matrix) && y >= 0 {
		// fmt.Println(x, y)
		cur := matrix[x][y]
		switch true {
		case cur == target:
			return true
		case cur > target:
			y--
		case cur < target:
			x++
		}
	}
	return false
}

// func main() {
// 	matrix := [][]int{
// 		[]int{1, 4, 7, 11, 15},
// 		[]int{2, 5, 8, 12, 19},
// 		[]int{3, 6, 9, 16, 22},
// 		[]int{10, 13, 14, 17, 24},
// 		[]int{18, 21, 23, 26, 30},
// 	}

// 	fmt.Println(findNumberIn2DArray(matrix, 5))
// 	fmt.Println(findNumberIn2DArray(matrix, 20))
// }
