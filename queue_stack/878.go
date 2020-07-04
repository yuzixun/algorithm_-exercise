package main

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	helper := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []rune{}
	for _, item := range s {
		_, ok := helper[item]
		if ok {
			if len(stack) > 0 && stack[len(stack)-1] == helper[item] {
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
