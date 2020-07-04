package main

func sortArray(nums []int) []int {
	// quickSort(nums, 0, len(nums)-1)
	// selectionSort(nums)
	// nums =
	// fmt.Println(nums)
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

// func selectionSort(nums []int) {
// 	minIndex := 0
// 	for i := 0; i < len(nums); i++ {
// 		minIndex = i

// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[minIndex] > nums[j] {
// 				minIndex = j
// 			}
// 		}

// 		nums[minIndex], nums[i] = nums[i], nums[minIndex]
// 	}

// 	fmt.Println(nums)
// }

// func quickSort(nums []int, left, right int) {
// 	if left < right {
// 		mid := partition(nums, left, right)
// 		sort(nums, left, mid-1)
// 		sort(nums, mid+1, right)
// 	}
// }

// func partition(nums []int, left, right int) int {
// 	q := left

// 	left++
// 	for left <= right {
// 		for left <= right && nums[left] <= nums[q] {
// 			left++
// 		}
// 		for left <= right && nums[right] > nums[q] {
// 			right--
// 		}
// 		if left > right {
// 			break
// 		}
// 		nums[left], nums[right] = nums[right], nums[left]
// 	}
// 	nums[right], nums[q] = nums[q], nums[right]

// 	return right
// }

func main() {
	sortArray([]int{5, 2, 3, 1})
	sortArray([]int{5, 1, 1, 2, 0, 0})
	sortArray([]int{0, 0, 1, 0, 0})
}
