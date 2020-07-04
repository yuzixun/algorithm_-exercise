package main

import "fmt"

// 逐个遍历 时间复杂度O(m+n)
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	len1, len2 := len(nums1), len(nums2)
// 	len := len1 + len2

// 	left, right := 0, 0
// 	index1, index2 := 0, 0

// 	for i := 0; i <= len/2; i++ {
// 		left = right
// 		if index1 < len1 && (index2 >= len2 || nums1[index1] < nums2[index2]) {
// 			right = nums1[index1]
// 			index1++
// 		} else {
// 			right = nums2[index2]
// 			index2++
// 		}
// 	}

// 	if len&1 == 0 {
// 		return float64(left+right) / 2.0
// 	}
// 	return float64(right)
// }
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size1, size2 := len(nums1), len(nums2)
	size := size1 + size2
	left := getKth(nums1, nums2, 0, size1-1, 0, size2-1, (size+1)/2)
	right := getKth(nums1, nums2, 0, size1-1, 0, size2-1, (size+2)/2)

	return (left + right) / 2.0
}

func getKth(nums1, nums2 []int, start1, end1, start2, end2, k int) float64 {
	size1, size2 := end1-start1+1, end2-start2+1
	if size1 > size2 {
		return getKth(nums2, nums1, start2, end2, start1, end1, k)
	}
	if size1 == 0 {
		return float64(nums2[start2+k-1])
	}
	if k == 1 {
		if nums1[start1] < nums2[start2] {
			return float64(nums1[start1])
		}
		return float64(nums2[start2])
	}

	newIndex1 := start1 + min(size1, k/2) - 1
	newIndex2 := start2 + min(size2, k/2) - 1

	if nums1[newIndex1] > nums2[newIndex2] {
		return getKth(nums1, nums2, start1, end1, newIndex2+1, end2, k-(newIndex2-start2+1))
	}
	return getKth(nums1, nums2, newIndex1+1, end1, start2, end2, k-(newIndex1-start1+1))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
