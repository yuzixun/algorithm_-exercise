package main

func surfaceArea(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			result = result + 2 + grid[i][j]*4
			if i-1 >= 0 {
				result = result - min(grid[i-1][j], grid[i][j])*2
			}
			if j-1 >= 0 {
				result = result - min(grid[i][j-1], grid[i][j])*2
			}
		}
	}
	// fmt.Println(result)
	return result
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	surfaceArea([][]int{[]int{2}})
	surfaceArea([][]int{[]int{}})
	surfaceArea([][]int{[]int{1, 2}, []int{3, 4}})
	surfaceArea([][]int{[]int{1, 0}, []int{0, 2}})
	surfaceArea([][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}})
	surfaceArea([][]int{[]int{2, 2, 2}, []int{2, 1, 2}, []int{2, 2, 2}})
}
