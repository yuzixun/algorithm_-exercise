package main

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count, result := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if result == nums[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			result = nums[i]
			count = 1
		}
	}
	return result
}

// func main() {
// 	fmt.Println(majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 3, 3, 3, 3}))
// }
