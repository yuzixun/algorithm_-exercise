package main

func reverseString(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {
	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
	reverseString([]byte{'H', 'a', 'n', 'n', 'a', 'h'})
}
