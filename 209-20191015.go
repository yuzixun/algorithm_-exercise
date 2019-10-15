package main

func minSubArrayLen(s int, nums []int) int {
	sum, size := 0, len(nums)
	min := size + 1

	for slow, fast := 0, 0; fast < size; fast++ {
		sum += nums[fast]

		for sum >= s {
			if min > fast-slow+1 {
				min = fast - slow + 1
			}
			sum -= nums[slow]
			slow++
		}
	}

	if min == size+1 {
		return 0
	}
	return min
}

// func main() {
// 	// s := 7
// 	// nums := []int{2, 3, 1, 2, 4, 3}
// 	s, nums := 11, []int{1, 2, 3, 4, 5}
// 	fmt.Println(minSubArrayLen(s, nums))
// }
