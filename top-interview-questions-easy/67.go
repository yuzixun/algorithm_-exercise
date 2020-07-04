package main

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}
	for i := 1; i < numRows; i++ {
		cur := make([]int, i+1)
		cur[0] = 1
		cur[i] = 1
		for j := 1; j < i; j++ {
			cur[j] = result[i-1][j-1] + result[i-1][j]
		}
		result[i] = cur
	}

	return result
}

func main() {
	generate(1)
	generate(5)
}
