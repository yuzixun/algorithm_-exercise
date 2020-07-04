package main

func reverseBits(num uint32) uint32 {
	arr := []uint32{}
	cur := uint32(1)
	for i := 0; i < 32; i++ {
		if num&cur == 0 {
			arr = append(arr, 0)
		} else {
			arr = append(arr, 1)
		}

		cur <<= 1
	}

	result := uint32(0)
	for i := 0; i < 32; i++ {
		result <<= 1
		result |= arr[i]
	}
	return result
}

func main() {
	reverseBits(00000010100101000001111010011100)
	reverseBits(11111111111111111111111111111101)
}
