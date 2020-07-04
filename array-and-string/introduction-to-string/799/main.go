package main

import "fmt"

func addBinary(a string, b string) string {
	sizeA, sizeB := len(a), len(b)
	if sizeB > sizeA {
		return addBinary(b, a)
	}

	carry := 0
	i, j := sizeA-1, sizeB-1

	result := make([]byte, sizeA)
	index := sizeA - 1
	for i >= 0 && j >= 0 {
		sum := int(a[i]-48) + int(b[j]-48) + carry
		result[index] = byte(sum%2 + 48)
		carry = sum / 2

		i--
		j--
		index--
	}

	for i >= 0 {
		sum := int(a[i]-48) + carry
		result[index] = byte(sum%2 + 48)
		carry = sum / 2

		i--
		index--
	}

	if carry == 1 {
		result = append([]byte{'1'}, result...)
	}
	return string(result)
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("11", "1111"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("0", "0"))
}
