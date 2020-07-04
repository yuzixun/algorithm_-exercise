package main

/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 */

func search(reader ArrayReader, target int) int {
	NULL := 2147483647
	left, right := 0, 20000

	for left <= right {
		mid := left + (right-left)/2

		cur := reader.get(mid)
		if cur > target {
			right = mid - 1
		} else if cur < target {
			left = mid + 1
		} else if cur == target {
			return mid
		}
	}

	return -1
}
