package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, i1, i2 := len(nums1)-1, m-1, n-1

	for i1 >= 0 && i2 >= 0 {
		if nums1[i1] > nums2[i2] {
			nums1[i] = nums1[i1]
			i1--
		} else {
			nums1[i] = nums2[i2]
			i2--
		}
		i--
	}

	for i2 >= 0 {
		nums1[i] = nums2[i2]
		i2--
		i--
	}

	// fmt.Println(nums1)
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
