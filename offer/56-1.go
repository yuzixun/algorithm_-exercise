package main

func singleNumbers(nums []int) []int {
	m := 0
	for i := 0; i < len(nums); i++ {
		m ^= nums[i]
	}

	min := 1
	for (min & m) != min {
		min <<= 1
	}

	a, b := 0, 0
	// withMin, withoutMin := []int{}, []int{}
	for i := 0; i < len(nums); i++ {
		if (min & nums[i]) == min {
			a ^= nums[i]
			// withMin = append(withMin, nums[i])
		} else {
			// withoutMin = append(withoutMin, nums[i])
			b ^= nums[i]
		}
	}

	return []int{a, b}
}

// func main() {
// 	fmt.Println(singleNumbers([]int{4, 1, 4, 6}))
// 	fmt.Println(singleNumbers([]int{1, 2, 10, 4, 1, 4, 3, 3}))
// 	fmt.Println(singleNumbers([]int{1, 2, 5, 2}))
// }
