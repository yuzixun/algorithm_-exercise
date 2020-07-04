package main

func reversePairs(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	copy := make([]int, size)
	for i, val := range nums {
		copy[i] = val
	}

	return reversePairsCore(&nums, &copy, 0, size-1)
}

func reversePairsCore(nums, copy *[]int, start, end int) int {
	if start == end {
		(*copy)[start] = (*nums)[start]
		return 0
	}

	size := (end - start) / 2

	left := reversePairsCore(copy, nums, start, start+size)
	// fmt.Println("left is ", nums, copy)
	right := reversePairsCore(copy, nums, start+size+1, end)
	// fmt.Println("right is ", nums, copy)

	i, j, copyIndex := start+size, end, end
	count := 0
	for i >= start && j >= start+size+1 {
		// fmt.Println((*nums)[i], (*nums)[j], start, end)
		if (*nums)[i] > (*nums)[j] {
			(*copy)[copyIndex] = (*nums)[i]
			// fmt.Println("value is ", (j - start - size))
			count += (j - start - size)
			i--
		} else {
			(*copy)[copyIndex] = (*nums)[j]
			j--
		}
		copyIndex--
	}

	for ; i >= start; i-- {
		(*copy)[copyIndex] = (*nums)[i]
		copyIndex--
	}

	for ; j >= start+size+1; j-- {
		(*copy)[copyIndex] = (*nums)[j]
		copyIndex--
	}

	return left + right + count
}

func main() {
	// fmt.Println(reversePairs([]int{7, 5, 6, 4}))
}
