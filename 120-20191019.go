package main

func minimumTotal(triangle [][]int) int {
	rows := len(triangle)
	result := make([]int, rows+1, rows+1)
	for i := rows - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if result[j] < result[j+1] {
				result[j] = result[j] + triangle[i][j]
			} else {
				result[j] = result[j+1] + triangle[i][j]
			}
			// fmt.Println("i is ", i, "result is ", result)
		}
	}
	// fmt.Println("result is ", result)
	return result[0]
}
