package main

import "fmt"

func splitArray(nums []int, m int) int {
	max, sum := getMaxAndSum(nums)
	if len(nums) == m {
		return max
	}
	left, right := max, sum
	var count, mid, tmp int
	for left < right {
		mid = left + (right-left)/2
		tmp = 0
		count = 1
		for _, num := range nums {
			tmp += num
			if tmp > mid {
				tmp = num
				count++
			}
		}
		if count > m {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func getMaxAndSum(nums []int) (int, int) {
	max, sum := -1, 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	return max, sum
}

func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
}
