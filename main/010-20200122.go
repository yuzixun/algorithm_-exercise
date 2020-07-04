package main

func isMatch1(s string, p string) bool {
	if p == "" {
		return s == ""
	}

	first := s != "" && (p[0] == s[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (first && isMatch(s[1:], p))
	}
	return first && isMatch(s[1:], p[1:])
}

// func main() {
// 	fmt.Println(isMatch("aa", "a"))
// 	fmt.Println(isMatch("aa", "a*"))
// 	fmt.Println(isMatch("ab", ".*"))
// 	fmt.Println(isMatch("aab", "c*a*b"))
// 	fmt.Println(isMatch("mississippi", "mis*is*p*."))
// }
