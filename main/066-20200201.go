package main

func plusOne(digits []int) []int {
	size := len(digits)

	carry := 1
	for i := size - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		if carry == 0 {
			break
		}
		digits[i] %= 10
	}
	if carry == 1 {
		a := []int{1}
		a = append(a, digits...)
		return a
	}
	return digits
}

// func main() {
// 	fmt.Println(plusOne([]int{1, 2, 3}))
// 	fmt.Println(plusOne([]int{4, 3, 2, 2}))
// 	fmt.Println(plusOne([]int{9}))
// 	fmt.Println(plusOne([]int{9, 9}))
// }
