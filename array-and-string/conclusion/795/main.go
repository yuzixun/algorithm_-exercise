package main

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return size
	}

	slow, fast := 0, 1
	for fast < size {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	// fmt.Println(slow, nums)
	return slow + 1
}

func main() {
	removeDuplicates([]int{1, 1, 2})
	removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
}
