package main

func isValid1(s string) bool {
	helper := map[rune]rune{')': '(', '}': '{', ']': '['}

	var stack []rune
	for _, i := range s {
		n := len(stack)
		if helper[i] != 0 {
			if n > 0 && helper[i] == stack[n-1] {
				stack = stack[0 : n-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, i)
		}
	}
	return len(stack) == 0
}

// func main() {
// 	fmt.Println(isValid1("()"))
// 	fmt.Println(isValid1("()[]{}"))
// 	fmt.Println(isValid1("(]"))
// 	fmt.Println(isValid1("([)]"))
// 	fmt.Println(isValid1("{[]}"))
// }
