package main

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	size := len(nums)
	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = 1
	}

	product := 1
	for i := 1; i < size; i++ {
		result[i], product = nums[i-1]*product, nums[i-1]*product
	}

	product = nums[size-1]
	for i := size - 2; i >= 0; i-- {
		result[i] = product * result[i]
		product = product * nums[i]
	}

	// fmt.Println(result)
	return result
}

func main() {
	productExceptSelf([]int{1, 2, 3, 4})
	productExceptSelf([]int{1, 2, 3, 4, 6, 91, 23, 4, 5})
	productExceptSelf([]int{0, 0})
}
