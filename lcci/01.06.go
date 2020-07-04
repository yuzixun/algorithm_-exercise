package main

import (
	"strconv"
)

func compressString(S string) string {
	if S == "" {
		return ""
	}
	cache := make([]byte, 0)
	count := 1
	prev := S[0]
	for i := 1; i < len(S); i++ {
		if prev != S[i] {
			cache = append(cache, prev)
			cache = append(cache, strconv.Itoa(count)...)
			count = 1
		} else {
			count++
		}
		prev = S[i]
	}
	cache = append(cache, prev)
	cache = append(cache, strconv.Itoa(count)...)
	if len(cache) < len(S) {
		return string(cache)
	}
	return S
}
