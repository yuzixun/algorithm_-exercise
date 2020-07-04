package main

// func isSubsequence(s string, t string) bool {
// 	if s == "" {
// 		return true
// 	}
// 	if t == "" {
// 		return false
// 	}
// 	return matchSubsequence(s, t, len(s)-1, len(t)-1)
// }

// func matchSubsequence(s, t string, sIndex, tIndex int) bool {
// 	if sIndex < 0 {
// 		return true
// 	}

// 	if tIndex < 0 {
// 		return false
// 	}

// 	if s[sIndex] == t[tIndex] {
// 		return matchSubsequence(s[0:sIndex], t[0:tIndex], sIndex-1, tIndex-1)
// 	}

// 	return matchSubsequence(s, t[0:tIndex], sIndex, tIndex-1)
// }
func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	if t == "" {
		return false
	}

	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

// func main() {
// 	fmt.Println(isSubsequence("abc", "ahbgdc"))
// 	fmt.Println(isSubsequence("axc", "ahbgdcx"))
// 	fmt.Println(isSubsequence("", "ahbgdcx"))
// 	fmt.Println(isSubsequence("b", "a"))
// 	fmt.Println(isSubsequence("abc", ""))
// }
