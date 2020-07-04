package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	var mid, result int

	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			result = 1
			tmpLeft, tmpRight := mid-1, mid+1
			for tmpLeft >= 0 && nums[tmpLeft] == target {
				// fmt.Println(tmpLeft, nums[tmpLeft], target, result)

				tmpLeft--
				result++
			}

			for tmpRight <= len(nums)-1 && nums[tmpRight] == target {
				// fmt.Println(tmpRight, nums[tmpRight], target, result)
				tmpRight++
				result++
			}
			// fmt.Println(tmpRight-tmpLeft+1, result)
			// break
			return tmpRight - tmpLeft + 1
		}
	}

	return 0
}

// func main() {
// 	search([]int{5, 7, 7, 8, 8, 10}, 8)
// 	search([]int{5, 7, 7, 8, 8, 10}, 6)
// }
