package main

func movingCount(m int, n int, k int) int {
	if m <= 0 || n <= 0 || k < 0 {
		return 0
	}

	visited := make(map[int]bool)

	return movingCountCore(m, n, k, 0, 0, visited)
}

func movingCountCore(m, n, k, i, j int, visited map[int]bool) int {
	var result int
	if checked(m, n, k, i, j, visited) {
		visited[i*n+j] = true
		result = 1 + movingCountCore(m, n, k, i-1, j, visited) +
			movingCountCore(m, n, k, i+1, j, visited) +
			movingCountCore(m, n, k, i, j-1, visited) +
			movingCountCore(m, n, k, i, j+1, visited)
	}
	return result
}

func checked(m, n, k, i, j int, visited map[int]bool) bool {
	return i >= 0 && j >= 0 && i < m && j < n && !visited[i*n+j] && calc(i, j, k)
}

func calc(i, j, k int) bool {
	sum := 0
	for i != 0 {
		sum += (i % 10)
		i /= 10
	}
	for j != 0 {
		sum += (j % 10)
		j /= 10
	}

	return sum <= k
}

// func main() {
// 	fmt.Println(movingCount(2, 3, 1))
// 	fmt.Println(movingCount(3, 1, 0))
// 	fmt.Println(movingCount(0, 0, 1))
// }
