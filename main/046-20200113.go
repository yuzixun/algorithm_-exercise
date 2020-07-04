package main

func dfs(cur, size int, nums, arr []int, token []bool, result *[][]int) {
	if cur == size {
		tmp := make([]int, size)
		copy(tmp, arr)
		*result = append(*result, tmp)
	}

	for i := 0; i < size; i++ {
		if !token[i] {
			token[i] = true
			arr[cur] = nums[i]
			dfs(cur+1, size, nums, arr, token, result)
			token[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	size := len(nums)
	arr, token := make([]int, size), make([]bool, size)

	var result [][]int
	dfs(0, size, nums, arr, token, &result)

	// fmt.Println(result)
	return result
}

// func main() {
// 	permute([]int{1, 2, 3})
// }
