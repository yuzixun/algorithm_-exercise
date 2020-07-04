package main

// func rotate(nums []int, k int) {
// 	size := len(nums)

// 	for i := 0; i < k; i++ {
// 		tmp := nums[size-1]
// 		for j := size - 1; j > 0; j-- {
// 			nums[j] = nums[j-1]
// 		}
// 		nums[0] = tmp
// 	}
// }

// func rotate(nums []int, k int) {
// 	k %= len(nums)
// 	swap(nums)
// 	swap(nums[:k])
// 	swap(nums[k:])
// 	// fmt.Println(nums, k)
// }

// func swap(nums []int) {
// 	size := len(nums)
// 	for i := 0; i < len(nums)/2; i++ {
// 		nums[i], nums[size-i-1] = nums[size-i-1], nums[i]
// 	}
// }

func rotate(nums []int, k int) {
	size := len(nums)
	k %= size
	count := 0
	for cur := 0; count < size; cur++ {
		tmp := cur
		pre := nums[cur]
		next := -1
		for cur != next {
			next = (tmp + k) % size
			nums[next], pre = pre, nums[next]
			tmp = next
			count++
		}
	}
	// fmt.Println(nums)
}
func main() {
	(rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3))
	(rotate([]int{-1, -100, 3, 99}, 2))
	(rotate([]int{-1}, 2))
}
