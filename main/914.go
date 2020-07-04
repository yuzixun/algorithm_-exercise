package main

import "fmt"

func hasGroupsSizeX(deck []int) bool {
	hash := map[int]int{}
	for _, val := range deck {
		hash[val]++
	}

	x := 0
	for _, val := range hash {
		x = gcd(x, val)
		if val < 2 {
			return false
		}
	}
	return x > 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
func main() {
	fmt.Println(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))
	fmt.Println(hasGroupsSizeX([]int{1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 2, 2, 2, 2}))
}
