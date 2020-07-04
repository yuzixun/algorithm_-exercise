package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	p, p1, p2 := m+n-1, m-1, n-1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] < nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
		} else {
			nums1[p] = nums1[p1]
			p1--
		}
		p--
	}

	for p1 >= 0 {
		nums1[p] = nums1[p1]
		p1--
		p--
	}
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
	// fmt.Println(nums1)
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	nums1 := []int{1, 2, 3, 0, 0, 0}
// 	nums2 := []int{2, 5, 6}

// 	merge(nums1, 3, nums2, 3)
// 	merge([]int{0}, 0, []int{1}, 1)
// }
