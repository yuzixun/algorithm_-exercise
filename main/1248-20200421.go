package main

func numberOfSubarrays(nums []int, k int) int {

	arr := make([]int, len(nums)+1)
	arr[0] = 0
	for index, num := range nums {
		arr[index+1] = arr[index] + num&1
	}
	hash := map[int]int{}
	result := 0
	// fmt.Println(arr)
	for _, item := range arr {
		result += hash[item-k]
		hash[item]++
	}
	// fmt.Println(hash, arr, result)
	return result
}

func main() {
	numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3)
	numberOfSubarrays([]int{2, 4, 6}, 1)
	numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 1, 2, 2}, 3)
}
