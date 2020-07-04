package main

func replaceSpace(s string) string {
	spaceCount := 0
	for _, i := range s {
		if i == ' ' {
			spaceCount++
		}
	}

	length := spaceCount*2 + len(s)
	result := make([]byte, length)
	space := "%20" // []rune{'%', '2', '0'}

	for i, j := length-1, len(s)-1; i >= 0; i-- {
		// fmt.Println(length, i, j)
		if s[j] == ' ' {
			result[i] = space[2]
			i--
			result[i] = space[1]
			i--
			result[i] = space[0]
		} else {
			result[i] = s[j]
		}
		j--
	}

	return string(result)
}

// func main() {
// 	fmt.Println(replaceSpace("We are happy."))
// 	// fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
// 	// fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
// }
