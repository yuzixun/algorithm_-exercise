package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	num, result := "", ""
	resStack := []string{}
	numStack := []string{}

	for _, item := range s {
		if item == '[' {
			resStack = append(resStack, result)
			numStack = append(numStack, num)

			num = ""
			result = ""
		} else if item == ']' {
			var n, res string
			n, numStack = numStack[len(numStack)-1], numStack[:len(numStack)-1]
			res, resStack = resStack[len(resStack)-1], resStack[:len(resStack)-1]
			count, _ := strconv.Atoi(n)
			result = res + strings.Repeat(result, count)
		} else if item >= '0' && item <= '9' {
			num += string(item)
		} else {
			result += string(item)
		}
	}

	// fmt.Println(result)
	return result
}

func main() {
	decodeString("3[a]2[bc]")
	decodeString("3[a2[c]]")
	decodeString("2[abc]3[cd]ef")
}
