package main

func hammingWeight(num uint32) int {
	i := 0
	for num != 0 {
		num &= (num - 1)
		i++
	}
	return i
}

// func main() {
// 	fmt.Println(hammingWeight(00000000000000000000000000001011))
// 	fmt.Println(hammingWeight(00000000000000000000000010000000))
// 	num := uint32(11111111111111111111111111111101)
// 	fmt.Println(hammingWeight(num))
// }
