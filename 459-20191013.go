package main

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	if s == "" {
		return false
	}
	size := len(s)
	ss := s[1:] + s[:size-1]
	// fmt.Println(ss)
	return strings.Contains(ss, s)
}
