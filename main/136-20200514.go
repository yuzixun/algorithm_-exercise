package main

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}

	// fmt.Println(result)
	return result
}

func main() {
	singleNumber([]int{2, 2, 1})
	singleNumber([]int{4, 1, 2, 1, 2})
}
