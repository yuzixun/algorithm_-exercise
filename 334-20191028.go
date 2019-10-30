package main

func increasingTriplet(nums []int) bool {
	a, b := 1<<32, 1<<32

	for _, num := range nums {
		switch {
		case a >= num:
			a = num
		case b >= num:
			b = num
		default:
			return true
		}
	}
	return false
}
