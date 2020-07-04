package main

func minSubArrayLen(s int, nums []int) int {
	size := len(nums)
	result := size + 1

	left, right, sum := 0, 0, 0
	for right < size {
		sum += nums[right]
		for sum >= s {
			if result > right-left+1 {
				result = right - left + 1
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	if result == size+1 {
		return 0
	}
	// fmt.Println(result)
	return result
}

func main() {
	minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	minSubArrayLen(9, []int{2, 3, 1, 2, 4, 3})
	minSubArrayLen(10, []int{2, 3, 1, 2, 4, 3})
	minSubArrayLen(10, []int{2, 3, 1, 2, 4, 3, 10})
	minSubArrayLen(6, []int{10, 2, 3})
	minSubArrayLen(3, []int{1, 1})
}
