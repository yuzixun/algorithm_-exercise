package main

func subarraySum(nums []int, k int) int {
	size := len(nums)
	hash := make(map[int]int, size)

	hash[0] = 1
	count := 0
	sum := 0
	for _, num := range nums {
		sum += num
		if hash[sum-k] > 0 {
			count += hash[sum-k]
		}

		hash[sum]++
	}

	// fmt.Println(count)
	return count
}

func main() {
	subarraySum([]int{3, 1, 1, 1, 1, 2, 2}, 2)
	subarraySum([]int{1}, 0)
	subarraySum([]int{1}, 2)
}
