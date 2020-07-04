package main

import "fmt"

func addBinary(a string, b string) string {
	aSize, bSize := len(a), len(b)
	if aSize < bSize {
		return addBinary(b, a)
	}

	result := make([]byte, aSize)
	aIndex, bIndex, rIndex := aSize-1, bSize-1, aSize-1
	carry := 0

	for aIndex >= 0 && bIndex >= 0 {
		cur := int(a[aIndex]-'0') + int(b[bIndex]-'0') + carry
		result[rIndex] = byte('0' + cur%2)
		carry = cur / 2

		aIndex--
		bIndex--
		rIndex--
	}

	for aIndex >= 0 {
		cur := int(a[aIndex]-'0') + carry
		result[rIndex] = byte('0' + cur%2)
		carry = cur / 2

		aIndex--
		rIndex--
	}

	if carry == 1 {
		result = append([]byte{'1'}, result...)
	}

	return string(result)
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}
