package main

import "fmt"

func waysToStep(n int) int {
	cache := make([]int, n+4)
	mod := 1000000007
	cache[0] = 0
	cache[1] = 1
	cache[2] = 2
	cache[3] = 4

	for i := 4; i <= n; i++ {
		cache[i] = (cache[i-1] + cache[i-2] + cache[i-3]) % mod
	}

	return cache[n]
}
func main() {
	fmt.Println(waysToStep(3))
	fmt.Println(waysToStep(5))
	fmt.Println(waysToStep(10000))
	fmt.Println(waysToStep(1000000))
}
