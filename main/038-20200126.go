package main

import (
	"strconv"
)

func countAndSay(n int) string {
	// hash := map[int]int{1: 1}
	// k := '1'
	// v := '1'

	// slow, fast := 0, 0
	bArr := []byte{'1'}
	counter := 1

	for i := 2; i <= n; i++ {
		cur := []byte{}
		for j := 0; j < len(bArr); j++ {
			if j == len(bArr)-1 || bArr[j] != bArr[j+1] {
				cur = append(cur, []byte(strconv.Itoa(counter))...)
				cur = append(cur, bArr[j])
				counter = 1
			} else {
				counter++
			}
			// fmt.Println(cur)
		}
		bArr = cur
	}
	// fmt.Println((string(bArr)))
	return string(bArr)
}

// func main() {
// 	countAndSay(1)
// 	// countAndSay(5)
// 	// countAndSay(3)
// 	// countAndSay(3)
// 	// countAndSay(3)
// 	// countAndSay(2)
// 	// countAndSay(1)
// }
