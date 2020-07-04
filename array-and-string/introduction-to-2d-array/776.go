package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	result := [][]int{
		[]int{1},
	}

	for i := 1; i < numRows; i++ {
		cur := make([]int, i+1)
		cur[0], cur[i] = 1, 1
		for j := 1; j <= i-1; j++ {
			cur[j] = result[i-1][j-1] + result[i-1][j]
		}
		result = append(result, cur)
	}

	// fmt.Println(result)
	return result
}

func main() {
	generate(5)
	generate(0)
}
