package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	helper := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, item := range s {
		v, ok := helper[item]
		if ok {
			if len(stack) > 0 && stack[len(stack)-1] == v {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, item)
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
