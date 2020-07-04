package main

func findKthLargest(nums []int, k int) int {
	for i := k / 2; i >= 0; i-- {
		minHeapify(nums[:k], i)
	}
	// fmt.Println(nums)

	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i]
			minHeapify(nums[:k], 0)
		}
	}
	// fmt.Println(nums)
	return nums[0]
}

func minHeapify(nums []int, root int) {
	left, right := 2*root+1, 2*root+2

	min := root
	if left < len(nums) && nums[min] > nums[left] {
		min = left
	}

	if right < len(nums) && nums[min] > nums[right] {
		min = right
	}

	if root != min {
		nums[min], nums[root] = nums[root], nums[min]
		minHeapify(nums, min)
	}

}

func main() {
	findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
}
