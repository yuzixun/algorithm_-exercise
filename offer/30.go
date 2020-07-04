package main

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
// func Constructor() MinStack {
// 	var stack []int
// 	var min []int
// 	return MinStack{stack: stack, min: min}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)

	length := len(this.min)

	if length == 0 || length > 0 && (this.min[length-1] >= x) {
		this.min = append(this.min, x)
	}

}

func (this *MinStack) Pop() {
	length := len(this.stack)

	result := this.stack[length-1]
	this.stack = this.stack[:length-1]

	if result == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {

	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

// func main() {
// 	obj := Constructor()
// 	obj.Push(0)
// 	obj.Push(1)
// 	obj.Push(0)
// 	// obj.Push(-2)
// 	// obj.Push(0)
// 	// obj.Push(-3)
// 	fmt.Println(obj.Min())
// 	obj.Pop()
// 	fmt.Println(obj.Min())
// }

// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.min();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.min();   --> 返回 -2.
