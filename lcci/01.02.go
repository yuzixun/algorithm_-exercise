package main

func CheckPermutation(s1 string, s2 string) bool {
	count := make(map[rune]int, len(s1))
	for _, item := range s1 {
		count[item]++
	}

	for _, item := range s2 {
		count[item]--
	}

	for _, value := range count {
		if value != 0 {
			return false
		}
	}

	return true
}
