package main

func findKthLargest(nums []int, k int) int {
	size := len(nums)
	k = size - k
	left, right := 0, size-1
	for left <= right {
		q := partition(nums, left, right)
		// fmt.Println("size is ", size, " q is ", q, " k is ", k, " left is ", left, " right is ", right)

		switch {
		case q > k:
			right = q - 1
		case q < k:
			left = q + 1
		default:
			return nums[q]
		}
	}

	return -1
}

func partition(nums []int, left, right int) int {
	// fmt.Println("left is ", left, "right is ", right, "nums is ", nums)
	x := nums[right]
	i := left

	for j := left; j <= right-1; j++ {
		if nums[j] < x {
			// fmt.Println(i, j, nums[j], x)
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[right] = nums[right], nums[i]
	return i
}

// func main() {
// 	// nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
// 	nums := []int{3, 2, 1, 5, 6, 4}
// 	k := 2
// 	// fmt.Println(findKthLargest(nums, k))
// }
