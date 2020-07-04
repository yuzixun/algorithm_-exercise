package main

func countPrimes(n int) int {
	cache := make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		cache[i] = true
	}
	count := 0
	for i := 2; i < n; i++ {
		if cache[i] {
			count++
			for j := 2; i*j < n; j++ {
				cache[i*j] = false
			}
		}
	}
	// fmt.Println(count)
	return count
}

func main() {
	countPrimes(10)
}
