package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[0:n]
}

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
	fmt.Println(reverseLeftWords("lrloseumgh", 6))
	fmt.Println(reverseLeftWords("av", 1))
}
