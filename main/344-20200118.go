package main

func reverseString(s []byte) {
	size := len(s)
	// fmt.Println(s)

	for i := 0; i < size/2; i++ {
		j := size - 1 - i
		s[i], s[j] = s[j], s[i]
	}
	// fmt.Println(s)
}

// func main() {
// 	reverseString([]byte{'h', 'e', 'l', 'x', 'o'})
// 	reverseString([]byte{'H', 'a', 'n', 'n', 'a', 'h'})
// }
