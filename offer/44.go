package main

import (
	"math"
)

func findNthDigit(n int) int {

	digits := 1
	for {
		counter := countOfInteger(digits)
		if n > counter*digits {
			n -= counter * digits
			digits++
			continue
		}

		return digitAtIndex(n, digits)
	}
	return -1
}

func countOfInteger(digits int) int {
	if digits == 1 {
		return 10
	}
	return 9 * int(math.Pow10(digits-1))
}

func digitAtIndex(index, digits int) int {
	number := beginNumber(digits) + index/digits
	indexFromRight := digits - index%digits
	for i := 1; i < indexFromRight; i++ {
		number /= 10
	}
	return number % 10
}

func beginNumber(digits int) int {
	if digits == 1 {
		return 0
	}
	return int(math.Pow10(digits - 1))
}

// func main() {
// 	fmt.Println(findNthDigit(3))
// 	fmt.Println(findNthDigit(11))
// 	fmt.Println(findNthDigit(0))
// }
