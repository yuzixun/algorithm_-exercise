package main

import "fmt"

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	return dfs(m, n, k, 0, 0, &visited)
}

func dfs(m, n, k, i, j int, visited *[][]bool) int {
	if i < 0 || i >= m || j < 0 || j >= n || getValue(i, j) > k || (*visited)[i][j] {
		return 0
	}

	(*visited)[i][j] = true

	return dfs(m, n, k, i-1, j, visited) + dfs(m, n, k, i, j-1, visited) + dfs(m, n, k, i, j+1, visited) + dfs(m, n, k, i+1, j, visited) + 1
}
func getValue(x, y int) int {
	sum := 0

	for x != 0 {
		sum += x % 10
		x = x / 10
	}

	for y != 0 {
		sum += y % 10
		y = y / 10
	}
	return sum
}
func main() {
	fmt.Println(movingCount(2, 3, 1))
	fmt.Println(movingCount(3, 1, 0))
}
