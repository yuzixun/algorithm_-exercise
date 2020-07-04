package main

import "fmt"

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	cache := make([]int, n+1)

	NULL := -1 << 32
	for i := n - 1; i >= 0; i-- {
		cache[i] = NULL
		s := 0
		for j := i; j < n && j < i+3; j++ {
			s += stoneValue[j]
			cache[i] = max(cache[i], s-cache[j+1])
		}
	}
	// fmt.Println("cache ", cache)
	if cache[0] > 0 {
		return "Alice"
	} else if cache[0] == 0 {
		return "Tie"
	} else {
		return "Bob"
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	// fmt.Println(stoneGameIII([]int{1, 2, 3, 7}))
	// fmt.Println(stoneGameIII([]int{1, 2, 3, -9}))
	fmt.Println(stoneGameIII([]int{1, 2, 3, 6}))
	// fmt.Println(stoneGameIII([]int{1, 2, 3, -1, -2, -3, 7}))
	fmt.Println(stoneGameIII([]int{-1, -2, -3}))
}
