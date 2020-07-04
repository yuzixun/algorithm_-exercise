package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rowLen := len(obstacleGrid[0])

	result := make([]int, rowLen)
	result[0] = 1
	for _, row := range obstacleGrid {
		for i := 0; i < len(row); i++ {
			if row[i] == 1 {
				result[i] = 0
			} else if i > 0 {
				result[i] += result[i-1]
			}
		}
	}
	return result[rowLen-1]
}

// func main() {
// 	obstacleGrid := [][]int{
// 		[]int{0, 0, 0},
// 		[]int{0, 1, 0},
// 		[]int{0, 0, 0},
// 	}
// 	// obstacleGrid := [][]int{
// 	// 	[]int{1},
// 	// }
// 	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
// }
