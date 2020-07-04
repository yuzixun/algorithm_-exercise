package main

func permute(nums []int) [][]int {
	res := [][]int{}
	size := len(nums)
	curArr := make([]int, size)
	visited := make([]bool, size)
	dfs(0, size, curArr, nums, visited, &res)
	return res
}

func dfs(curIndex, size int, curArr, nums []int, visited []bool, res *[][]int) {
	if curIndex == size {
		tmp := make([]int, size)
		copy(tmp, curArr)
		(*res) = append((*res), tmp)
	}

	for i := 0; i < size; i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		curArr[curIndex] = nums[i]
		dfs(curIndex+1, size, curArr, nums, visited, res)
		visited[i] = false
	}
}

func main() {
	permute([]int{1, 2, 3})
}
