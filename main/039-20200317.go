package main

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	current := []int{}
	dfs(candidates, 0, target, current, &result)
	return result
}

func dfs(candidates []int, begin, target int, current []int, result *[][]int) {
	if target == 0 {
		r := make([]int, len(current))
		copy(r, current)
		(*result) = append((*result), r)
		return
	} else if target < 0 {
		return
	}

	for i := begin; i < len(candidates); i++ {
		if target < candidates[i] {
			continue
		}
		current = append(current, candidates[i])
		dfs(candidates, i, target-candidates[i], current, result)
		current = current[:len(current)-1]
	}
}
