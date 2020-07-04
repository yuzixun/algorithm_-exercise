package main

import "fmt"

func minWindow(s string, t string) string {
	left, right, valid := 0, 0, 0
	need, cur := map[byte]int{}, map[byte]int{}
	start, size := 0, 1<<32

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	for right < len(s) {
		rItem := s[right]
		right++

		if _, ok := need[rItem]; ok {
			cur[rItem]++
			if cur[rItem] == need[rItem] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < size {
				start = left
				size = right - left
			}

			lItem := s[left]
			left++

			if _, ok := need[lItem]; ok {
				if cur[lItem] == need[lItem] {
					valid--
				}
				cur[lItem]--
			}
		}
	}

	if size == 1<<32 {
		return ""
	}
	return s[start : start+size]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
