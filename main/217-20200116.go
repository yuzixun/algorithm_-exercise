package main

func containsDuplicate(nums []int) bool {
	hash := make(map[int]bool, len(nums))
	for _, num := range nums {
		if hash[num] {
			return true
		}
		hash[num] = true
	}
	return false
}

// func main() {
// 	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
// 	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
// 	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
// }
