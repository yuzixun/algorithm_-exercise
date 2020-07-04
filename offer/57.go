package main

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left <= right {
		cur := nums[left] + nums[right]
		if cur > target {
			right--
		} else if cur < target {
			left++
		} else {
			break
		}
	}
	return []int{nums[left], nums[right]}
}

// func main() {
// 	twoSum([]int{2, 7, 11, 15}, 9)
// }
