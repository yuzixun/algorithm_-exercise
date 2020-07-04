package main

func uniquePaths(m int, n int) int {
	var cache [][]int

	for i := 0; i < m; i++ {
		l := make([]int, n)

		for j := 0; j < n; j++ {
			l[j] = 1
		}

		cache = append(cache, l)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cache[i][j] = cache[i-1][j] + cache[i][j-1]
		}
	}
	return cache[m-1][n-1]
}

// func calculate(m, n int, cache *[][]int) int {
// 	if m < 1 || n < 1 {
// 		return 0
// 	}

// 	if m == 1 || n == 1 {
// 		return 1
// 	}

// 	if (*cache)[m-1][n-1] != 0 {
// 		return (*cache)[m-1][n-1]
// 	}

// 	result := calculate(m-1, n, cache) + calculate(m, n-1, cache)
// 	(*cache)[m-1][n-1] = result

// 	return result
// }

// func main() {
// 	fmt.Println(uniquePaths(3, 2))
// 	fmt.Println(uniquePaths(7, 3))
// }
