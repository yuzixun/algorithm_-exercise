package main

// func convert(s string, numRows int) string {
// 	if numRows == 1 {
// 		return s
// 	}

// 	resArr := make([]bytes.Buffer, numRows)
// 	flag, current, first := 1, -1, true

// 	for i := 0; i < len(s); i++ {

// 		if flag > 0 {
// 			current++
// 		} else {
// 			current--
// 		}
// 		resArr[current].WriteByte(s[i])
// 		if current == numRows-1 || (current == 0 && !first) {
// 			flag *= -1
// 		}
// 		if first {
// 			first = false
// 		}
// 	}

// 	result := make([]string, numRows)
// 	for i := 0; i < numRows; i++ {
// 		result[i] = resArr[i].String()
// 	}

// 	return strings.Join(result, "")
// }

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	resArr := make([][]byte, numRows)
	// resArr := make([]bytes.Buffer, numRows)
	flag, current := false, 0

	for i := 0; i < len(s); i++ {
		resArr[current] = append(resArr[current], s[i])

		if current == numRows-1 || current == 0 {
			flag = !flag
		}

		if flag {
			current++
		} else {
			current--
		}
	}

	var result string
	for i := 0; i < numRows; i++ {
		result+= string(resArr[i])
	}

	return result
}

// func main() {
// 	fmt.Println(convert("LEETCODEISHIRING", 3)) // LCIRETOESIIGEDHN
// 	fmt.Println(convert("LEETCODEISHIRING", 4)) // LDREOEIIECIHNTSG
// 	fmt.Println(convert("LEETCODEISHIRING", 1)) // LDREOEIIECIHNTSG
// }
