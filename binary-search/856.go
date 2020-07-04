package main

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	existMap := map[int]int{}

	for _, num := range nums1 {
		existMap[num]++
	}

	for _, num := range nums2 {
		if existMap[num] > 0 {
			existMap[num]--
			result = append(result, num)
		}
	}
	// fmt.Println(result)
	return result
}

func main() {
	intersection([]int{1, 2, 2, 1}, []int{2, 2})
	intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
}
