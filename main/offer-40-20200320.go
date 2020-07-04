package main

import "fmt"

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}
	current, start, end := 0, 0, len(arr)-1
	for {
		current = partition(arr, start, end)
		if current == k {
			return arr[:k]
		} else if current < k {
			start = current + 1
			current = partition(arr, start, end)
		} else {
			end = current - 1
			current = partition(arr, start, end)
		}
	}

}

func partition(arr []int, start, end int) int {
	q := (arr)[start]
	left, right := start+1, end
	for left <= right {
		for left < end && (arr)[left] <= q {
			left++
		}

		for right > start && (arr)[right] >= q {
			right--
		}

		if left >= right {
			break
		}

		(arr)[left], (arr)[right] = (arr)[right], (arr)[left]
		left++
		right--
	}

	(arr)[start], (arr)[right] = (arr)[right], (arr)[start]
	return right
}

func main() {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
	fmt.Println(getLeastNumbers([]int{0, 1, 2, 1}, 1))
}
