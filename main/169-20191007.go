package main

func majorityElement(nums []int) int {
	majMap := make(map[int]int)

	for _, num := range nums {
		majMap[num]++
	}

	var maxCount, maxNum int
	for num, count := range majMap {
		if maxCount < count {
			maxCount = count
			maxNum = num
		}
	}

	return maxNum
}
