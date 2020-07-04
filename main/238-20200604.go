package main

func productExceptSelf(nums []int) []int {
	size := len(nums)
	if size == 0 {
		return []int{}
	}
	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = 1
	}

	product := 1
	for i := 0; i < size; i++ {
		result[i] = product
		product *= nums[i]
	}

	product = 1
	for i := size - 1; i >= 0; i-- {
		result[i] *= product
		product *= nums[i]
	}

	// fmt.Println(result)
	return result
}

func main() {
	productExceptSelf([]int{1, 2, 3, 4, 5})
}
