package main

func moveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}
	for ; count < len(nums); count++ {
		nums[count] = 0
	}
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0, 1, 3, 12})
	moveZeroes([]int{1, 3, 12})
	moveZeroes([]int{1})
	moveZeroes([]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0})
}
