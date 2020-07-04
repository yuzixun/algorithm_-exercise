package main

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result[i-1] = "FizzBuzz"
		} else if i%5 == 0 {
			result[i-1] = "Buzz"
		} else if i%3 == 0 {
			result[i-1] = "Fizz"
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}
	// fmt.Println(result)
	return result
}

func main() {
	fizzBuzz(15)
}
