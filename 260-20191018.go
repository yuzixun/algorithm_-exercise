package main

func singleNumber(nums []int) []int {
	tmp := 0
	for _, num := range nums {
		tmp ^= num
	}

	r := tmp & -tmp

	var r1, r2 int
	for _, num := range nums {
		if num&r == r {
			r1 ^= num
		} else {
			r2 ^= num
		}
	}

	return []int{r1, r2}
}

// func main() {
// 	nums := []int{1, 2, 1, 3, 2, 5}
// 	singleNumber(nums)
// }
