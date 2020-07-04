package main

func findMaxConsecutiveOnes(nums []int) int {
	result := 0

	slow := 0
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] != 1 {
			cur := i - slow
			if cur > result {
				result = cur
			}
			slow = i + 1
		}
	}
	if i-slow > result {
		result = i - slow
	}

	// fmt.Println(result)
	return result
}

func main() {
	findMaxConsecutiveOnes([]int{1, 1, 0, 0, 0, 0, 1, 1, 1})
}
