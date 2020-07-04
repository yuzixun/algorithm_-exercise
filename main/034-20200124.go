package main

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}

	findLeft, findRight := true, false
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if findLeft {
				if mid == 0 || nums[mid-1] != target {
					findLeft, findRight = false, true
					result[0] = mid
					left, right = 0, len(nums)-1
				} else if nums[mid-1] == target {
					right = mid - 1
				}
			}

			if findRight {
				if mid == len(nums)-1 || nums[mid+1] != target {
					result[1] = mid
					break
				} else if nums[mid+1] == target {
					left = mid + 1
				}
			}

		}
	}
	// fmt.Println(left, right, left+(right-left)/2)
	// fmt.Println(result)
	return result
}

// func main() {
// 	searchRange([]int{8, 8, 8, 8}, 8)
// 	searchRange([]int{5, 7, 7, 8, 8, 8, 10}, 8)
// 	searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
// }
