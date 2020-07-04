package main

func constructArr(a []int) []int {
	product := 1
	size := len(a)
	if size == 0 {
		return []int{}
	}
	result := make([]int, size)
	result[0] = 1

	for i := 1; i < size; i++ {
		product *= a[i-1]
		result[i] = product
	}

	product = 1
	for i := size - 2; i >= 0; i-- {
		product *= a[i+1]
		result[i] *= product
	}

	return result
}

func main() {
	constructArr([]int{1, 2, 3, 4, 5})
}
