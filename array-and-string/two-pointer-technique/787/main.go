package main

func removeElement(nums []int, val int) int {
	size := len(nums)

	slow := 0
	for i := 0; i < size; i++ {
		if nums[i] != val {
			nums[slow] = nums[i]
			slow++
		}
	}

	// fmt.Println(nums)
	return slow + 1
}

func main() {
	removeElement([]int{3, 2, 2, 3}, 3)
	removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
}
