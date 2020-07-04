package main

func longestConsecutive(nums []int) int {
	hash := make(map[int]int)
	cur, maxLength := 0, 0
	for _, num := range nums {
		if hash[num] != 0 {
			continue
		}
		left, right := hash[num-1], hash[num+1]

		cur = left + right + 1
		if cur > maxLength {
			maxLength = cur
		}

		hash[num] = cur
		hash[num-left] = cur
		hash[num+right] = cur
		// fmt.Println(hash)
	}
	// fmt.Println(maxLength)
	return maxLength
}

func main() {
	longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	longestConsecutive([]int{})
	longestConsecutive([]int{1, 2, 0, 1})
}
