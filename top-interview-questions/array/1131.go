package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums, cur []int
}

func Constructor(nums []int) Solution {
	s := Solution{nums: nums}
	s.cur = make([]int, len(nums))
	copy(s.cur, nums)
	fmt.Println(s)
	return s
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	copy(this.cur, this.nums)
	return this.cur
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := 0; i < len(this.nums); i++ {
		j := i + rand.Intn(len(this.nums)-i)
		this.cur[i], this.cur[j] = this.cur[j], this.cur[i]
	}
	return this.cur
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func main() {

}
