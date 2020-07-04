package main

type MinStack struct {
	stack    []int
	minStack []int
	size     int
	minSize  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
		size:     0,
		minSize:  0,
	}

}

func (this *MinStack) Push(x int) {
	if this.minSize == 0 || this.minStack[this.minSize-1] >= x {
		this.minStack = append(this.minStack, x)
		this.minSize++
	}

	this.stack = append(this.stack, x)
	this.size++
	// fmt.Println(this.stack, this.minStack)
}

func (this *MinStack) Pop() {
	cur := this.stack[this.size-1]
	this.stack = this.stack[:this.size-1]
	this.size--

	if cur == this.minStack[this.minSize-1] {
		this.minStack = this.minStack[:this.minSize-1]
		this.minSize--
	}
}

func (this *MinStack) Top() int {
	return this.stack[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[this.minSize-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
}
