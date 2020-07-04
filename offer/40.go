package main

func partition(arr []int, start, end int) int {
	q := arr[start]
	left, right := start+1, end

	for left <= right {
		for left < end && arr[left] <= q {
			left++
		}

		for right > start && arr[right] >= q {
			right--
		}

		if left >= right {
			break
		}
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	arr[start], arr[right] = arr[right], arr[start]
	return right
}

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}

	left, right := 0, len(arr)-1
	index := partition(arr, left, right)
	for index != k {
		if index < k {
			left = index + 1
			index = partition(arr, left, right)
		} else if index > k {
			right = index - 1
			index = partition(arr, left, right)
		}
	}

	// fmt.Println(arr)
	return arr[:k]
}

// func main() {
// 	getLeastNumbers([]int{3, 2, 1}, 2)
// 	getLeastNumbers([]int{0, 1, 2, 1}, 1)
// 	getLeastNumbers([]int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}, 8)
// }
