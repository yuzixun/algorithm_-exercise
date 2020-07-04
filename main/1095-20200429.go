package main

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	left, right := 0, mountainArr.length()-1
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if mountainArr.get(mid) < mountainArr.get(mid+1) {
			left = mid + 1
		} else {
			right = mid
		}
	}

	maxIndex := left
	left, right = 0, maxIndex

	for left <= right {
		mid = left + (right-left)/2
		value := mountainArr.get(mid)
		if value == target {
			return mid
		} else if value > target {
			right = mid - 1
		} else if value < target {
			left = mid + 1
		}
	}

	left, right = maxIndex, mountainArr.length()-1
	for left <= right {
		mid = left + (right-left)/2
		value := mountainArr.get(mid)
		if value == target {
			return mid
		} else if value > target {
			left = mid + 1
		} else if value < target {
			right = mid - 1
		}
	}

	return -1
}
