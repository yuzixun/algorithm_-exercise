package main

import "fmt"

var cnt = [2]int{}

func patternMatching(pattern string, value string) bool {
	pSize, vSize := len(pattern), len(value)
	// 如果模式为空，字符串也为空，则匹配
	if pSize == 0 {
		return vSize == 0
	}

	// 如果字符串为空，则模式串必须都是由相同字符组成的
	if vSize == 0 {
		i := 0
		for i < pSize && pattern[i] == pattern[0] {
			i++
		}

		return i == pSize
	}

	cnt[0], cnt[1] = 0, 0
	for _, x := range pattern {
		cnt[x-'a']++
	}

	// 其中一个长度为0，则判断另一个，是否匹配
	if cnt[0] == 0 {
		return helper(value, cnt[1])
	} else if cnt[1] == 0 {
		return helper(value, cnt[0])
	}

	// 假设其中一个未空字符串，与另一个是否匹配
	if helper(value, cnt[0]) || helper(value, cnt[1]) {
		return true
	}

	for i := 1; i*cnt[0] <= vSize-cnt[1]; i++ {
		if (vSize-i*cnt[0])%cnt[1] != 0 {
			continue
		}

		j := (vSize - i*cnt[0]) / cnt[1]
		if check(pattern, value, i, j) {
			return true
		}
	}

	return false
}

func helper(value string, k int) bool {
	vSize := len(value)
	if vSize%k != 0 {
		return false
	}

	size := vSize / k
	for i := size; i < vSize; i += size {
		if value[i:i+size] != value[0:size] {
			return false
		}
	}

	return true
}

func check(pattern, value string, aSize, bSize int) bool {
	ps := []string{"", ""}
	for i, j := 0, 0; i < len(pattern); i++ {
		if pattern[i] == 'a' {
			if ps[0] == "" {
				ps[0] = value[j : j+aSize]
			} else if value[j:j+aSize] != ps[0] {
				return false
			}

			j += aSize

		} else if pattern[i] == 'b' {

			if ps[1] == "" {
				ps[1] = value[j : j+bSize]
			} else if value[j:j+bSize] != ps[1] {
				return false
			}

			j += bSize
		}
	}

	return ps[0] != ps[1]
}

func main() {
	fmt.Println(patternMatching("abba", "dogcatcatdog"))
	fmt.Println(patternMatching("abba", "dogcatcatfish"))
	fmt.Println(patternMatching("aaaa", "dogcatcatdog"))
	fmt.Println(patternMatching("abba", "dogdogdogdog"))
}
