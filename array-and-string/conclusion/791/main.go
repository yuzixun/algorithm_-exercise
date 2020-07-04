package main

func rotate(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])

	// fmt.Println(nums)
}

func reverse(nums []int) {
	size := len(nums)
	for i := 0; i < size/2; i++ {
		nums[i], nums[size-i-1] = nums[size-i-1], nums[i]
	}
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotate([]int{-1, -100, 3, 99}, 2)
}
