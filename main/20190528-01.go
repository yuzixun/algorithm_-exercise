package main

func twoSum(nums []int, target int) []int {

	numToIndex := make(map[int]int)
	for i, num := range nums {
		_, d := numToIndex[num]
		if d {
			return []int{numToIndex[num], i}
		} else {
			numToIndex[target-num] = i
		}
	}

	return nil
}

// func twoSum(nums []int, target int) []int {
// 	result := []int{}
// 	len := len(nums)
// 	for i := 0; i < len; i++ {
// 		num := nums[i]
// 		for index := i + 1; index < len; index++ {
// 			if num+nums[index] == target {
// 				result = []int{i, index}
// 				break
// 			}
// 		}
// 	}
// 	return result
// }

// func main() {
// 	nums := []int{2, 7, 11, 15}
// 	target := 26

// 	result := twoSum(nums, target)

// 	fmt.Println(result)
// }
