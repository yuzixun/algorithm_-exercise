package main

func add(a int, b int) int {
	for b != 0 {
		t := a ^ b
		b = (a & b) << 1
		a = t
	}

	return a
}

func main() {
	add(1, 1)
	add(9, 2)
	add(13, 120)
}
