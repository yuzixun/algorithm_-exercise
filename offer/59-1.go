package main

func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	if len(nums) == 0 {
		return result
	}

	current := nums[0]
	for i := 1; i < k; i++ {

		current = max(current, nums[i])

	}
	result = append(result, current)

	slow, fast := 1, k
	for fast < len(nums) {
		// fmt.Println(slow, fast, slow+k-1, result[slow-1])
		if nums[slow-1] == result[slow-1] {
			current = nums[slow]
			for i := slow + 1; i <= fast; i++ {
				current = max(current, nums[i])
			}
			result = append(result, current)
		} else if nums[fast] > result[slow-1] {
			result = append(result, nums[fast])
		} else {
			result = append(result, result[len(result)-1])
		}
		slow++
		fast++
	}
	// fmt.Println(result)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
}
