package main

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	mid := 0
	for numbers[left] >= numbers[right] {
		if left+1 == right {
			mid = right
			break
		}

		mid = left + (right-left)/2
		if numbers[mid] == numbers[left] && numbers[mid] == numbers[right] {
			return nFind(numbers, left, right)
		}

		if numbers[left] <= numbers[mid] {
			left = mid
		} else if numbers[mid] <= numbers[right] {
			right = mid
		}
		// fmt.Println(left, mid, right, numbers[left], numbers[mid], numbers[right])
	}
	// fmt.Println(mid)
	return numbers[mid]
}

func nFind(numbers []int, left, right int) int {
	var result int
	for i := left; i < right; i++ {
		if numbers[i] > result {
			result = numbers[i]
		}
	}
	return result
}

// func main() {
// 	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
// 	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
// 	fmt.Println(minArray([]int{4, 4, 4, 0, 1, 2, 3}))
// 	fmt.Println(minArray([]int{1, 2, 3}))
// 	fmt.Println(minArray([]int{1}))
// 	fmt.Println(minArray([]int{1, 1, 1, 0, 1}))
// 	fmt.Println(minArray([]int{4, 4, 4, 4, 0, 1, 2, 3}))
// }
