package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	h *IntHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	var kthLargest KthLargest
	if len(nums) <= k {
		kthLargest = KthLargest{h: (*IntHeap)(&nums), k: k}
		heap.Init(kthLargest.h)

	} else {
		firstKNums := nums[:k]
		kthLargest = KthLargest{h: (*IntHeap)(&firstKNums), k: k}
		heap.Init(kthLargest.h)
		for _, num := range nums[k:] {
			kthLargest.Add(num)
		}

	}
	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	if this.h.Len() < this.k {
		heap.Push(this.h, val)
	} else if val > (*this.h)[0] {
		(*this.h)[0] = val
		heap.Fix(this.h, 0)
	}
	return (*this.h)[0]
}

func main() {
	k := 3
	arr := []int{4, 5, 8, 2}
	obj := Constructor(k, arr)
	fmt.Println(obj)
	for _, i := range arr {
		param := obj.Add(i)
		fmt.Println("result is ", param)
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
