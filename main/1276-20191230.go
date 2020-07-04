package main

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	a := float64(tomatoSlices-2*cheeseSlices) / 2
	b := float64(4*cheeseSlices-tomatoSlices) / 2

	modA := (tomatoSlices - 2*cheeseSlices) % 2
	modB := (4*cheeseSlices - tomatoSlices) % 2
	// fmt.Println(modA, modB)
	if (a < 0 || b < 0) || (modA != 0 || modB != 0) {
		return []int{}
	}
	return []int{int(a), int(b)}
}

// func main() {
// 	fmt.Println(numOfBurgers(16, 7))
// 	fmt.Println(numOfBurgers(17, 4))
// 	fmt.Println(numOfBurgers(4, 17))
// 	fmt.Println(numOfBurgers(0, 0))
// 	fmt.Println(numOfBurgers(2, 1))
// }
