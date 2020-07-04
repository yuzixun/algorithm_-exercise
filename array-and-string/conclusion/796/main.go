package main

func moveZeroes(nums []int) {
	size := len(nums)
	if size <= 1 {
		return
	}

	slow, fast := 0, 0
	for ; fast < size; fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}

	for i := slow; i < size; i++ {
		nums[i] = 0
	}
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}
