package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	allSize := len(nums1) - 1
	size1, size2 := m-1, n-1

	for size1 >= 0 && size2 >= 0 {
		if nums1[size1] > nums2[size2] {
			nums1[allSize] = nums1[size1]
			size1--
			allSize--
		} else {
			nums1[allSize] = nums2[size2]
			allSize--
			size2--
		}
	}

	for size2 >= 0 {
		nums1[allSize] = nums2[size2]
		allSize--
		size2--
	}
	// fmt.Println(nums1)
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{0, 0, 0}, 0, []int{2, 5, 6}, 3)
}
