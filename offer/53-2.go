package main

// func missingNumber(nums []int) int {
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] != i {
// 			return i
// 		}
// 	}
// 	return nums[len(nums)-1] + 1
// }

func missingNumber(nums []int) int {

	left, right := 0, len(nums)-1
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// fmt.Println(left, nums[left])
	if left == nums[left] {
		return left + 1
	}
	return left
}

func main() {
	missingNumber([]int{0, 1, 3})
	missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9})
	missingNumber([]int{0})
}
