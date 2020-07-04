package main

func hammingWeight(num uint32) int {
	var count int
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	// fmt.Println(count)
	return count
}

// func main() {
// 	hammingWeight(00000000000000000000000000001011)
// 	hammingWeight(00000000000000000000000010000000)
// 	// hammingWeight(11111111111111111111111111111101)
// }
