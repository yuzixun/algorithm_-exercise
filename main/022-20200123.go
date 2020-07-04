package main

func generateParenthesis(n int) []string {
	var result []string
	if n == 0 {
		return result
	}
	dfs("", n, n, &result)
	return result
}

func dfs(cur string, left, right int, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, cur)
		return
	}

	if left > right {
		return
	}

	if left > 0 {
		dfs(cur+"(", left-1, right, result)
	}

	if right > 0 {
		dfs(cur+")", left, right-1, result)
	}
}

// func main() {
// 	generateParenthesis(3)
// }
