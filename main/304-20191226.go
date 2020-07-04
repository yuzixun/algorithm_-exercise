package main

type NumMatrix struct {
	memo [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		return NumMatrix{}
	}
	n := len(matrix[0])

	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}

	memo[0][0] = matrix[0][0]
	for i := 1; i < m; i++ {
		memo[i][0] = matrix[i][0] + memo[i-1][0]
	}

	for j := 1; j < n; j++ {
		memo[0][j] = matrix[0][j] + memo[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			memo[i][j] = memo[i-1][j] + memo[i][j-1] - memo[i-1][j-1] + matrix[i][j]
		}
	}
	return NumMatrix{memo: memo}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	result := this.memo[row2][col2]
	if row1 > 0 {
		result -= this.memo[row1-1][col2]
	}
	if col1 > 0 {
		result -= this.memo[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		result += this.memo[row1-1][col1-1]
	}
	return result
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

// func main() {
// 	matrix := [][]int{
// 		[]int{3, 0, 1, 4, 2},
// 		[]int{5, 6, 3, 2, 1},
// 		[]int{1, 2, 0, 1, 5},
// 		[]int{4, 1, 0, 1, 7},
// 		[]int{1, 0, 3, 0, 5},
// 	}
// 	nm := Constructor(matrix)
// 	fmt.Println(nm.SumRegion(0, 0, 1, 1)) // -> 8
// 	fmt.Println(nm.SumRegion(2, 1, 4, 3)) // -> 8
// 	fmt.Println(nm.SumRegion(1, 1, 2, 2)) // -> 11
// 	fmt.Println(nm.SumRegion(1, 2, 2, 4)) // -> 12
// }
