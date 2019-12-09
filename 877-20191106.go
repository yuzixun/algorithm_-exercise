package main

import "fmt"

func stoneGame(piles []int) bool {
	size := len(piles)
	// left, right := 0, size-1

	result := make([][][]int, size, size)
	for j := range result {
		result[j] = make([][]int, size, size)
		for i := range result[j] {
			result[j][i] = []int{0, 0}
		}
	}

	for i := 0; i < size; i++ {
		result[i][i][0] = piles[i]
	}

	for i := 2; i <= size; i++ {
		for j := 0; j <= size-1; j++ {
			k := i + j - 1

			if k > size-1 {
				break
			}
			// fmt.Println(j, k)
			left := piles[j] + result[j+1][k][1]
			right := piles[k] + result[j][k-1][1]

			if left > right {
				result[j][k][0] = left
				result[j][k][1] = result[j+1][k][0]
			} else {
				result[j][k][0] = right
				result[j][k][1] = result[j][k-1][0]
			}
		}
	}
	// fmt.Println(result)
	item := result[0][size-1]
	return item[0] > item[1]
}

func main() {
	// piles := []int{5, 3, 4, 5}
	piles := []int{3, 2, 10, 4}
	fmt.Println(stoneGame(piles))
}
