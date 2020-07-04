package main

func singleNumber2(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		a = (a ^ num) & ^b
		b = (b ^ num) & ^a
	}
	return a
}

// func main() {
// 	nums := []int{2, 2, 3, 2}
// 	singleNumber(nums)
// }
