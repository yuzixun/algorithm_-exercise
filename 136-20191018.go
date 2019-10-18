package main

func singleNumber1(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// func main() {
// 	nums := []int{4, 1, 2, 1, 2}
// 	fmt.Println("singleNumber ", singleNumber(nums))
// }
