package main

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	result, current := [][]int{}, []int{}
	sort.Ints(candidates)
	dfs(candidates, 0, target, current, &result)
	// fmt.Println(result)
	return result
}

func dfs(candidates []int, begin, target int, current []int, res *[][]int) {
	// fmt.Println(begin, target, current)
	if target == 0 {
		r := make([]int, len(current))
		copy(r, current)
		(*res) = append((*res), r)
		return
	}
	if len(candidates) == begin {
		return
	}

	for i := begin; i < len(candidates); i++ {
		if i != begin && candidates[i] == candidates[i-1] {
			continue
		}
		if target < candidates[i] {
			return
		}
		current = append(current, candidates[i])
		dfs(candidates, i+1, target-candidates[i], current, res)
		current = current[:len(current)-1]
	}
}

func main() {
	combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	combinationSum2([]int{2, 5, 2, 1, 2}, 5)
}
