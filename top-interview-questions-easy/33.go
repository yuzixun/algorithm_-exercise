package main

import "fmt"

func reverse(x int) int {
	MAX := 1<<31 - 1
	MIN := -1 << 31
	result := 0
	for x != 0 {
		result = result*10 + x%10
		// fmt.Println("r is ", result)
		x = x / 10
	}

	if result > MAX || result < MIN {
		return 0
	}
	return result
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(-1))
	fmt.Println(reverse(1534236469))
}
