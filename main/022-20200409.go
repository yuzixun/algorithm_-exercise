package main

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	res := []string{}
	dfs(n, n, "", &res)
	return res
}

func dfs(left, right int, cur string, res *[]string) {
	if left > right {
		return
	}

	if left == 0 && right == 0 {
		*res = append(*res, cur)
	}

	if left > 0 {
		dfs(left-1, right, cur+"(", res)
	}
	if right > 0 {
		dfs(left, right-1, cur+")", res)
	}
}

func main() {
	generateParenthesis(3)
	generateParenthesis(4)
}
