package main

func minIncrementForUnique(A []int) int {
	arr := make([]int, 40001)
	minValue, maxValue := 40001, 0
	for _, num := range A {
		arr[num]++
		minValue = min(num, minValue)
		maxValue = max(num, maxValue)
	}

	dup := 0
	result := 0
	for minValue < maxValue {
		if arr[minValue] > 1 {
			dup = arr[minValue] - 1
			result += dup
			arr[minValue] = 1
			arr[minValue+1] += dup
		}
		minValue++
	}
	// fmt.Println("maxValue is ", maxValue)
	dup = arr[maxValue] - 1
	result += dup * (dup + 1) / 2
	// fmt.Println(result)
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	minIncrementForUnique([]int{1, 2, 2})
	minIncrementForUnique([]int{3, 2, 1, 2, 1, 7})
	minIncrementForUnique([]int{81, 7, 64, 45, 23, 45, 15, 33, 65, 47, 55, 54, 82, 34, 81, 7, 42, 35, 45, 83, 51, 98, 86, 11, 31, 2, 73, 1, 88, 19, 41, 6, 31, 38, 5, 76, 71, 35, 83, 7, 5, 35, 90, 75, 67, 56, 80, 41, 3, 80, 19, 38, 35, 12, 35, 17, 61, 65, 1, 45, 55, 87, 49, 8, 5, 86, 79, 85, 26, 64, 9, 82, 33, 95, 47, 3, 90, 62, 28, 3, 37, 15, 54, 14, 17, 92, 64, 24, 87, 62, 94, 48, 88, 67, 4, 24, 28, 17, 23, 22})
}
