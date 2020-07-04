package main

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	cur := 0

	for cur <= right {
		switch nums[cur] {
		case 0:
			nums[left], nums[cur] = nums[cur], nums[left]
			left++
			cur++
		case 2:
			nums[right], nums[cur] = nums[cur], nums[right]
			right--
		case 1:
			cur++
		}
	}
}

// func main() {
// 	sortColors([]int{2, 0, 2, 1, 1, 0})
// }
