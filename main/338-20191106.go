package main

func countBits(num int) []int {
	result := make([]int, num+1, num+1)

	for i := 0; i <= num; i++ {
		result[i] = result[i/2] + (i & 1)
	}

	return result
}

// func main() {
// 	fmt.Println(countBits(2))
// 	fmt.Println(countBits(5))
// 	fmt.Println(countBits(7))
// }
