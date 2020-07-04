package main

func distributeCandies(candies int, num_people int) []int {
	result := make([]int, num_people)
	current, curIndex := 0, 0
	for {
		curIndex = current % num_people
		current++
		if current > candies {
			result[curIndex] += candies
			break
		} else {
			result[curIndex] += current
			candies -= current
		}
	}
	return result
}

// func main() {
// 	fmt.Println(distributeCandies(7, 4))
// 	fmt.Println(distributeCandies(10, 3))
// }
