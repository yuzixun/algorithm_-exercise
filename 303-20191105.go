package main

import "fmt"

type NumArray struct {
	result []int
}

func Constructor(nums []int) NumArray {
	numArray := new(NumArray)
	size := len(nums)
	if size == 0 {
		return *numArray
	}

	numArray.result = make([]int, size, size)
	numArray.result[0] = nums[0]

	for i := 1; i < size; i++ {
		numArray.result[i] = numArray.result[i-1] + nums[i]
	}
	fmt.Println(numArray)
	return *numArray
}

func (this *NumArray) SumRange(i int, j int) int {
	tmp := 0
	if i-1 > 0 {
		tmp = this.result[i-1]
	}
	return this.result[j] - tmp
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	fmt.Println(obj.SumRange(1, 3))
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
	fmt.Println(obj.SumRange(4, 5))

}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
