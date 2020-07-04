package main

func sumNums(n int) int {
	var helper func(res *int, n int) bool
	helper = func(res *int, n int) bool {
		(*res) += n
		return n != 0 && helper(res, n-1)
	}
	res := 0
	helper(&res, n)
	return res
}

func main() {
	sumNums(3)
	sumNums(9)
}
