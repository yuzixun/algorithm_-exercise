package main

func generateMatrix(n int) [][]int {
	left, right, up, down, index := 0, n-1, 0, n-1, 1
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	for {
		for i := left; i <= right; i++ {
			result[up][i] = index
			index++
		}
		up++

		if index > n*n {
			break
		}

		for i := up; i <= down; i++ {
			result[i][right] = index
			index++
		}
		right--

		if index > n*n {
			break
		}

		for i := right; i >= left; i-- {
			result[down][i] = index
			index++
		}
		down--

		if index > n*n {
			break
		}
		for i := down; i >= up; i-- {
			result[i][left] = index
			index++
		}
		left++
		if index > n*n {
			break
		}
	}

	// fmt.Println(result)
	return result
}

// func main() {
// 	generateMatrix(1)
// 	generateMatrix(2)
// 	generateMatrix(3)
// }
