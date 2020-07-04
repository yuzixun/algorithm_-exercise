package main

func mySqrt(x int) int {
	result := x
	for result*result > x {
		result = (result + x/result) / 2
	}
	// fmt.Println(result)
	return result
}

func main() {
	mySqrt(4)
	mySqrt(8)
	mySqrt(9)
}
