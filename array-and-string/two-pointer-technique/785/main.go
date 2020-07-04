package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	var cur int
	for left < right {
		cur = numbers[left] + numbers[right]
		if cur == target {
			return []int{left + 1, right + 1}
		} else if cur > target {
			right--
		} else if cur < target {
			left++
		}
	}

	return []int{0, 0}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
