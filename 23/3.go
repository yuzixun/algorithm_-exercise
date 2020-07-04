package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 int) {

}
func checkOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {

	radiusXMin, radiusXMax := x_center-radius, x_center+radius
	radiusYMin, radiusYMax := y_center-radius, y_center+radius

	centerX, centerY := float32(x1+x2)/2, float32(y1+y2)/2

	deltaX, deltaY := math.Abs(centerX-float32(x_center)), math.Abs(centerY-float32(y_center))

	if deltaX <= math.Abs(centerX-radiusXMin) || deltaX <= math.Abs(centerX-radiusXMax) || deltaY <= math.Abs(centerX-radiusXMin) || deltaY <= math.Abs(centerX-radiusXMax) {

	}
	fmt.Println(radiusXMin, radiusXMax, radiusYMin, radiusYMax, radiusXMax < x1, radiusXMin > x2, radiusYMax < y1, radiusYMin > y2)
	if (radiusXMax < x1 || radiusXMin > x2) || (radiusYMax < y1 || radiusYMin > y2) {
		return false
	}
	return true
}

func main() {
	// fmt.Println(checkOverlap(1, 0, 0, 1, -1, 3, 1))
	// fmt.Println(checkOverlap(1, 0, 0, -1, 0, 0, 1))
	// fmt.Println(checkOverlap(1, 1, 1, -3, -3, 3, 3))
	// fmt.Println(checkOverlap(1, 1, 1, 1, -3, 2, -1))
	checkOverlap(1415, 807, -784, -733, 623, -533, 1005)

}
