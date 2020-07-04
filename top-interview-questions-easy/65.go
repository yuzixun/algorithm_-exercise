package main

func hammingDistance(x int, y int) int {
	num := x ^ y

	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}

	// fmt.Println(count)
	return count
}

func main() {
	hammingDistance(1, 4)
	hammingDistance(23, 100)
}
