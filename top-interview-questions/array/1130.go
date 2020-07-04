package main

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}

	for ; slow < len(nums); slow++ {
		nums[slow] = 0
	}
	// fmt.Println(nums)
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0, 1, 0, 0, 3, 12, 0})
	moveZeroes([]int{0, 1})
}
