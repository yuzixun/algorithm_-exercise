package main

import "fmt"

func mySqrt(x int) int {
	result := x
	for result*result > x {
		result = (result + x/result) / 2
	}
	return result
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(180))
}
