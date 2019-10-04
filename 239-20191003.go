package main

func maxSlidingWindow(nums []int, k int) []int {
	size := len(nums)
	if size == 0 {
		return []int{}
	}

	indexes := make([]int, k, k)
	results := make([]int, 0, size+k-1)

	for i, value := range nums {
		if i >= k && indexes[0] <= i-k {
			indexes = indexes[1:]
		}
		for len(indexes) > 0 && nums[indexes[len(indexes)-1]] <= value {
			indexes = indexes[:len(indexes)-1]
		}
		indexes = append(indexes, i)
		if i >= k-1 {
			results = append(results, nums[indexes[0]])
		}
	}
	return results
}

// func main() {
// 	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
// 	k := 3

// 	fmt.Println("result is ", maxSlidingWindow(nums, k))
// }
