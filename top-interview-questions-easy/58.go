package main

import (
	"math/rand"
)

type Solution struct {
	nums, current []int
	size          int
}

func Constructor(nums []int) Solution {
	size := len(nums)
	current := make([]int, size)
	copy(current, nums)
	return Solution{
		nums:    nums,
		current: current,
		size:    size,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := 0; i < this.size; i++ {
		j := i + rand.Intn(this.size-i)
		this.current[i], this.current[j] = this.current[j], this.current[i]
	}
	return this.current
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func main() {

}
