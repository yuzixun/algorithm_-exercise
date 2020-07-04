package main

func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	end := make([]int, size)
	end[0] = nums[0]
	right := 0
	result := 1

	for i := 1; i < size; i++ {
		l := 0
		r := right
		for l <= r {
			m := (l + r) / 2
			if nums[i] > end[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		}

		right = max(right, l)
		end[l] = nums[i]
		result = max(result, l+1)
	}

	// fmt.Println(end, result)
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
// 	lengthOfLIS(nums)
// }
