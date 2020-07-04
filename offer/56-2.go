package main

func singleNumber(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		a = (a ^ num) & ^b
		b = (b ^ num) & ^a
	}
	return a
	// var count, moved, result int
	// for i := 0; i < 32; i++ {
	// 	count = 0
	// 	moved = (1 << uint(i))
	// 	for _, num := range nums {
	// 		if num&moved != 0 {
	// 			count++
	// 		}
	// 	}

	// 	// fmt.Println("count is ", count)
	// 	if count%3 == 1 {
	// 		result |= moved
	// 	}
	// }
	// return result
}

// func main() {
// 	fmt.Println(singleNumber([]int{3, 4, 3, 3}))
// 	fmt.Println(singleNumber([]int{9, 1, 7, 9, 7, 9, 7}))
// }
