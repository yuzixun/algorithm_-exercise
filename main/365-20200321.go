package main

func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}

	if x == 0 || y == 0 {
		return z == 0 || x+y == z
	}

	for y != 0 {
		x, y = y, x%y
	}

	return z%x == 0
}
