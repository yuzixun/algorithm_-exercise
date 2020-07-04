package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size1, size2 := len(nums1), len(nums2)
	left := helper(nums1, nums2, 0, size1-1, 0, size2-1, (size1+size2+1)/2)
	right := helper(nums1, nums2, 0, size1-1, 0, size2-1, (size1+size2+2)/2)
	return float64(left+right) / 2.0
}

func helper(nums1, nums2 []int, left1, right1, left2, right2, k int) float64 {
	size1, size2 := right1-left1+1, right2-left2+1
	if size2 > size1 {
		return helper(nums2, nums1, left2, right2, left1, right1, k)
	}

	if size2 == 0 {
		return float64(nums1[left1+k-1])
	}

	if k == 1 {
		return float64(min(nums1[left1], nums2[left2]))
	}

	pos1 := left1 + min(k/2, size1) - 1
	pos2 := left2 + min(k/2, size2) - 1
	if nums1[pos1] < nums2[pos2] {
		return helper(nums1, nums2, pos1+1, right1, left2, right2, k-min(k/2, size1))
	}
	return helper(nums1, nums2, left1, right1, pos2+1, right2, k-min(k/2, size2))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
