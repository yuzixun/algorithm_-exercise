package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	symbols := []string{"+", "-", "*", "/"}

	for i := 0; i < len(tokens); i++ {
		isSym := false
		for _, symbol := range symbols {
			if symbol == tokens[i] {
				isSym = true
			}
		}

		if isSym {
			v1, v2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch tokens[i] {
			case "*":
				stack = append(stack, v1*v2)
			case "/":
				stack = append(stack, v1/v2)
			case "+":
				stack = append(stack, v1+v2)
			case "-":
				stack = append(stack, v1-v2)
			}
		} else {
			v, err := strconv.Atoi(tokens[i])
			if err == nil {
				stack = append(stack, v)
			}
		}
	}
	fmt.Println(stack)
	return stack[0]
}

// func main() {
// 	evalRPN([]string{"2", "1", "-"})
// 	evalRPN([]string{"2", "1", "+"})
// 	evalRPN([]string{"2", "1", "+", "3", "*"})
// 	evalRPN([]string{"4", "13", "5", "/", "+"})
// 	evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
// }
