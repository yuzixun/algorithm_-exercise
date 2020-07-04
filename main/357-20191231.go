package main

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	pre, cur := 10, 10
	for i := 2; i <= n; i++ {
		cur = pre + 9*mathA(i-1)
		pre = cur
	}

	return cur
}

func mathA(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= (9 - i)
	}
	return result
}

// func main() {
// 	fmt.Println(countNumbersWithUniqueDigits(0))
// 	fmt.Println(countNumbersWithUniqueDigits(1))
// 	fmt.Println(countNumbersWithUniqueDigits(2))
// 	fmt.Println(countNumbersWithUniqueDigits(3))
// 	fmt.Println(countNumbersWithUniqueDigits(4))
// 	fmt.Println(countNumbersWithUniqueDigits(5))
// 	fmt.Println(countNumbersWithUniqueDigits(6))
// }
