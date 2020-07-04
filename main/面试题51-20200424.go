package main

import "fmt"

func reversePairs(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	newNums := make([]int, size)
	copy(newNums, nums)

	count := 0
	var merge func(nums, newNums []int, start, mid, end int)
	var mergeSort func(nums, newNums []int, start, end int)

	merge = func(nums, newNums []int, start, mid, end int) {
		i, j, newIndex := start, mid+1, start

		for i <= mid && j <= end {
			if nums[i] <= nums[j] {
				newNums[newIndex] = nums[i]
				i++
			} else {
				count += (mid - (i - 1))

				newNums[newIndex] = nums[j]
				j++
			}
			newIndex++
		}

		for i <= mid {
			newNums[newIndex] = nums[i]
			i++

			newIndex++
		}

		for j <= end {
			newNums[newIndex] = nums[j]
			j++

			newIndex++
		}

		for i := start; i <= end; i++ {
			nums[i] = newNums[i]
		}
	}

	mergeSort = func(nums, newNums []int, start, end int) {
		if start >= end {
			return
		}
		mid := start + (end-start)/2
		mergeSort(nums, newNums, start, mid)
		mergeSort(nums, newNums, mid+1, end)
		merge(nums, newNums, start, mid, end)
	}

	mergeSort(nums, newNums, 0, len(nums)-1)
	return count
}

func main() {
	fmt.Println(reversePairs([]int{7, 5, 6, 4}))
}
