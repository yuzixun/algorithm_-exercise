package main

func mySqrt(x int) int {
	start, end, t := 1, x, 0
	for start <= end {
		t = start + (end-start)/2
		c := x / t
		if t == c {
			return t
		} else if t > c {
			end = t - 1
		} else if t < c {
			start = t + 1
		}
	}

	return end
}

// func main() {
// 	n := 1000
// 	fmt.Println(mySqrt(n))
// }
