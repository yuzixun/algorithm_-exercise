package main

func lastRemaining(n int, m int) int {
	last := 0

	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}

	return last
}

func main() {
	lastRemaining(5, 3)
	lastRemaining(10, 17)
}
