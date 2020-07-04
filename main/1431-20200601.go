package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	max := 0
	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}

	for i, candy := range candies {
		if candy+extraCandies >= max {
			result[i] = true
		}
	}

	// fmt.Println(result)
	return result
}

func main() {
	kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	kidsWithCandies([]int{4, 2, 1, 1, 2}, 1)
	kidsWithCandies([]int{4, 2, 1, 1, 2}, 10)
	kidsWithCandies([]int{12, 1, 12}, 10)
	kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
}
