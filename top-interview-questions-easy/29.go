package main

func twoSum(nums []int, target int) []int {

	hash := make(map[int]int, len(nums))

	for i, num := range nums {
		cur := target - num
		j, ok := hash[cur]
		if ok {
			return []int{i, j}
		}
		hash[cur] = i
	}

	return []int{}
}

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}
